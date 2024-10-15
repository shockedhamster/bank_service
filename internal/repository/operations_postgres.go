package repository

import (
	"fmt"

	"github.com/bank_service/internal/entities"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type OperationsPostgres struct {
	db *sqlx.DB
}

const (
	dateHighToLow   = "date_high_to_low"
	dateLowToHigh   = "date_low_to_high"
	amountHighToLow = "amount_high_to_low"
	amountLowToHigh = "amount_low_to_high"
)

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

func (r *OperationsPostgres) GetTransactionHistoryById(id int, sortType string, limit, offset int) ([]entities.Operation, error) {
	var resultTransactions entities.Operation
	var resultTransactionList []entities.Operation
	var orderBy string
	switch sortType {
	case dateHighToLow:
		orderBy = "o.created DESC"
	case dateLowToHigh:
		orderBy = "o.created"
	case amountHighToLow:
		orderBy = "o.amount DESC"
	case amountLowToHigh:
		orderBy = "o.amount"
	}

	query := fmt.Sprintf("select o.id, o.account_id, o.amount, ot.type_name, o.created from %s o join %s ot ON ot.id = o.operation_type where o.account_id = $1 Order by %s LIMIT $2 OFFSET $3", operationsTable, operationTypeTable, orderBy)
	rows, err := r.db.Query(query, id, limit, offset)
	if err != nil {
		logrus.Errorf("Error getting transaction history: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&resultTransactions.Id, &resultTransactions.AccountId, &resultTransactions.Amount, &resultTransactions.OperationType, &resultTransactions.Created)
		if err != nil {
			logrus.Errorf("Error scanning a row: %s", err.Error())
			return nil, err
		}

		resultTransactionList = append(resultTransactionList, resultTransactions)
	}

	return resultTransactionList, nil

}
