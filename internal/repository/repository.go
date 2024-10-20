package repository

import (
	"github.com/bank_service/internal/entities"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Account
	Operations
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Account:       NewAccountsPostgres(db),
		Operations:    NewOperationsPostgres(db),
	}
}

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GetUser(username, password string) (entities.User, error)
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
