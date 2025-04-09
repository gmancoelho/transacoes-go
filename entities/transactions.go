package entities

import (
	"github.com/gmancoelho/transacoes-go/formatter"
)

type Transactions struct {
	Value    int    `json:"valor"`
	DateHour string `json:"dataHora"`
}

func NewTransaction(value int, dateHour string) *Transactions {
	formattedDate, err := formatter.ConvertISO8601ToString(dateHour)
	if err != nil {
		return nil
	}
	return &Transactions{
		Value:    value,
		DateHour: formattedDate,
	}
}
