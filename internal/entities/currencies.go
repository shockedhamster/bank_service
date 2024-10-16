package entities

type Meta struct {
	LastUpdatedAt string `json:"last_updated_at"`
}

type Currency struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

type Data struct {
	USD Currency `json:"USD,omitempty"`
	EUR Currency `json:"EUR,omitempty"`
}

type CurrencyResp struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

type CurrencyReq struct {
	Currencies    string `json:"currencies"`
	Base_currency string `json:"base_currency"`
}
