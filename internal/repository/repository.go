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

// TODO: Deposit, Withdraw, Transfer
type Account interface {
	Deposit(id, amount int) error
	Withdraw(id, amount int) error
	//Transfer()
}

// TODO: GetUserBalanceById, GetTransactionHistoryById
type Operations interface {
	GetUserBalanceById(id int) (int, error)
}
