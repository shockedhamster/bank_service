package entities

import "time"

type Operation struct {
	Id            int       `json:"id" db:"id"`
	AccountId     int       `json:"account_id" db:"account_id"`
	Amount        int       `json:"amount" db:"amount"`
	OperationType string    `json:"operation_type" db:"operation_type"`
	Created       time.Time `json:"created" db:"created"`
}

const (
	OperationTypeDeposit      = "deposit"
	OperationTypeWithdraw     = "withdraw"
	OperationTypeTransferFrom = "transfer_from"
	OperationTypeTransferTo   = "transfer_to"
)
