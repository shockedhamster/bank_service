package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type OperationsPostgres struct {
	db *sqlx.DB
}

func NewOperationsPostgres(db *sqlx.DB) *OperationsPostgres {
	return &OperationsPostgres{db: db}
}

// Доработать для разных валют
func (r *OperationsPostgres) GetUserBalanceById(id int) (int, error) {
	var balance int
	query := fmt.Sprintf("SELECT balance FROM %s JOIN users ON users.id = accounts.id WHERE users.id = $1", accoutsTable)
	err := r.db.Get(&balance, query, id)
	if err != nil {
		logrus.Errorf("Error getting a balance: %s", err.Error())
		return -1, err
	}

	return balance, nil
}
