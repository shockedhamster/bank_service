package service

import (
	"github.com/bank_service/internal/entities"
	"github.com/bank_service/internal/repository"
)

type Service struct {
	Authorization
	Account
	Operations
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		Account:       NewAccountService(repository.Account),
		Operations:    NewOperationsService(repository.Operations),
	}
}

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Account interface {
	Deposit(id, amount int) error
	Withdraw(id, amount int) error
	Transfer(idFrom, idTo, amount int) error
}

type Operations interface {
	GetUserBalanceById(id int) (int, error)
	GetTransactionHistoryById(id int, sortType string, limit, offset int) ([]entities.Operation, error)
}
