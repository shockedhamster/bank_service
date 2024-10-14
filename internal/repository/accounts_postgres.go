package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AccountsPotgres struct {
	db *sqlx.DB
}

func NewAccountsPostgres(db *sqlx.DB) *AccountsPotgres {
	return &AccountsPotgres{db: db}
}

func (r *AccountsPotgres) Deposit(id, amount int) error {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("Error starting tx in Deposit: %s", err.Error())
		return err
	}
	insertIntoAccountsQuery := fmt.Sprintf("UPDATE %s SET balance = balance + %d WHERE id = %d", accoutsTable, amount, id)
	_, err = tx.Exec(insertIntoAccountsQuery)
	if err != nil {
		tx.Rollback()
		return err

	}

	insertIntoOperationsQuery := fmt.Sprintf("INSERT INTO %s (account_id, amount, operation_type) VALUES ($1, $2, (SELECT id FROM operation_type WHERE type_name='deposit'))", operationsTable)
	_, err = tx.Exec(insertIntoOperationsQuery, id, amount)
	if err != nil {
		tx.Rollback()
		return err

	}

	tx.Commit()
	return nil
}

func (r *AccountsPotgres) Withdraw(id, amount int) error {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("Error starting tx in withdraw: %s", err.Error())
		return err
	}
	insertIntoAccountsQuery := fmt.Sprintf("UPDATE %s SET balance = balance - %d WHERE id = %d", accoutsTable, amount, id)
	_, err = tx.Exec(insertIntoAccountsQuery)
	if err != nil {
		tx.Rollback()
		return err

	}

	insertIntoOperationsQuery := fmt.Sprintf("INSERT INTO %s (account_id, amount, operation_type) VALUES ($1, $2, (SELECT id FROM operation_type WHERE type_name='withdraw'))", operationsTable)
	_, err = tx.Exec(insertIntoOperationsQuery, id, amount)
	if err != nil {
		tx.Rollback()
		return err

	}

	tx.Commit()
	return nil
}

// func (r *AccountsPotgres) Transfer() {

// }
