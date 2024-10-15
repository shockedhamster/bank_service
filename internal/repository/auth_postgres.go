package repository

import (
	"fmt"

	"github.com/bank_service/internal/entities"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entities.User) (int, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("Error starting tx in transfer: %s", err.Error())
		return 0, err
	}
	createUserQuery := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id", usersTable)

	createUserRow := r.db.QueryRow(createUserQuery, user.Username, user.Password)

	if err := createUserRow.Scan(&id); err != nil {
		logrus.Errorf("Error creating a user: %s", err.Error())
		tx.Rollback()
		return 0, err
	}

	createAccountQuery := fmt.Sprintf("INSERT INTO %s (id, balance) values ($1, 0)", accoutsTable)
	_, err = r.db.Exec(createAccountQuery, id)
	if err != nil {
		logrus.Errorf("Error creating an account: %s", err.Error())
		tx.Rollback()
		return 0, err

	}
	tx.Commit()
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (entities.User, error) {
	var user entities.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	if err != nil {
		logrus.Errorf("Error getting a user: %s", err.Error())
	}
	return user, nil
}
