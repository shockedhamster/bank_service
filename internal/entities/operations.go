package entities

import "time"

type Operation struct {
	Id            int       `json:"id" db:"id"`
	AccountId     int       `json:"account_id" db:"account_id"`
	Amount        int       `json:"amount" db:"amount"`
	OperationType string    `json:"operation_type" db:"operation_type"`
	Created       time.Time `json:"created" db:"created"`
}

type GetTransactionHistoryInput struct {
	Id       int    `json:"id" db:"id"`
	SortType string `json:"sort_type" db:"sort_type"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}
