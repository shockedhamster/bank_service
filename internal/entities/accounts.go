package entities

type Account struct {
	Id      int `json:"id" db:"id" binding:"required"`
	Balance int `json:"balance" db:"balance" binding:"required"`
}

type Transfer struct {
	IdFrom int `json:"id_from" db:"id_from"`
	IdTo   int `json:"id_to" db:"id_to"`
	Amount int `json:"amount" db:"amount"`
}
