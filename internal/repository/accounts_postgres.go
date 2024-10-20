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

func (r *AccountsPotgres) CreateAccount(id int) error {
	createAccountQuery := fmt.Sprintf("INSERT INTO %s (id, balance) values ($1, 0)", accoutsTable)
	_, err := r.db.Exec(createAccountQuery, id)
	if err != nil {
		logrus.Errorf("Error creating an account: %s", err.Error())
		return err

	}
	return nil
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

func (r *AccountsPotgres) Transfer(idFrom, idTo, amount int) error {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("Error starting tx in transfer: %s", err.Error())
		return err
	}

	decreaseBalanceQuery := fmt.Sprintf("UPDATE %s SET balance = balance - %d WHERE id = %d", accoutsTable, amount, idFrom)
	_, err = tx.Exec(decreaseBalanceQuery)
	if err != nil {
		tx.Rollback()
		return err

	}
	increaseBalanceQuery := fmt.Sprintf("UPDATE %s SET balance = balance + %d WHERE id = %d", accoutsTable, amount, idTo)
	_, err = tx.Exec(increaseBalanceQuery)
	if err != nil {
		tx.Rollback()
		return err

	}

	insertIntoOperationsFromQuery := fmt.Sprintf("INSERT INTO %s (account_id, amount, operation_type) VALUES ($1, $2, (SELECT id FROM operation_type WHERE type_name='transfer_from'))", operationsTable)
	_, err = tx.Exec(insertIntoOperationsFromQuery, idFrom, amount)
	if err != nil {
		tx.Rollback()
		return err

	}

	insertIntoOperationsToQuery := fmt.Sprintf("INSERT INTO %s (account_id, amount, operation_type) VALUES ($1, $2, (SELECT id FROM operation_type WHERE type_name='transfer_to'))", operationsTable)
	_, err = tx.Exec(insertIntoOperationsToQuery, idTo, amount)
	if err != nil {
		tx.Rollback()
		return err

	}

	tx.Commit()
	return nil
}
