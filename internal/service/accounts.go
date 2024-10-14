package service

import (
	"github.com/bank_service/internal/repository"
	"github.com/sirupsen/logrus"
)

type AccoutsService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccoutsService {
	return &AccoutsService{
		repo: repo,
	}
}

func (s *AccoutsService) Deposit(id, amount int) error {
	err := s.repo.Deposit(id, amount)
	if err != nil {
		logrus.Errorf("Error making deposit: %s", err.Error())
		return err
	}
	return nil
}

func (s *AccoutsService) Withdraw(id, amount int) error {
	err := s.repo.Withdraw(id, amount)
	if err != nil {
		logrus.Errorf("Error making withdraw: %s", err.Error())
		return err
	}
	return nil
}

// func (s *AccoutsService) Transfer() {

// }
