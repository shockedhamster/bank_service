package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/bank_service/internal/entities"
	"github.com/bank_service/internal/kafka"
	"github.com/bank_service/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

const (
	salt       = "sadf23qfl/zxcv"
	signingKey = "sdjkallnfal322fr2g2"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo     repository.Authorization
	producer kafka.Producer
}

func NewAuthService(repo repository.Authorization, producer kafka.Producer) *AuthService {
	return &AuthService{
		repo:     repo,
		producer: producer,
	}
}

func (s *AuthService) SendMessage(topic string, key, message string) error {

	err := s.producer.ProduceMessage(topic, key, message)
	if err != nil {
		logrus.Errorf("Error sending message: %s", err.Error())
		return err
	}
	return nil

}

func (s *AuthService) CreateUser(user entities.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}
	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
