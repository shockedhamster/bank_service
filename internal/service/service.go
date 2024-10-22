package service

import (
	"github.com/bank_service/internal/entities"
	"github.com/bank_service/internal/kafka"
	"github.com/bank_service/internal/repository"
)

type Service struct {
	Authorization
	Account
	Operations
}

func NewService(repository *repository.Repository, producer kafka.Producer) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization, producer),
		Account:       NewAccountService(repository.Account),
		Operations:    NewOperationsService(repository.Operations),
	}
}

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	SendMessage(topic string, key, message string) error
}

type Account interface {
	CreateAccount(id int) error
	Deposit(id, amount int) error
	Withdraw(id, amount int) error
	Transfer(idFrom, idTo, amount int) error
}

type Operations interface {
	GetUserBalanceById(id int) (int, error)
	GetTransactionHistoryById(id int, sortType string, limit, offset int) ([]entities.Operation, error)
}
