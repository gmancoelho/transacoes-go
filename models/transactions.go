package models

type Transactions struct {
	Value    int    `json:"valor"`
	DateHour string `json:"dataHora"`
}

func NewTransaction(value int, dateHour string) *Transactions {
	return &Transactions{
		Value:    value,
		DateHour: dateHour,
	}
}
