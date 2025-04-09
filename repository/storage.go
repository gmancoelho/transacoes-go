package storage

import (
	e "github.com/gmancoelho/transacoes-go/entities"
)

type Storage interface {
	CreateTransaction(*e.Transactions) error
	DeleteTransaction(int) error
	GetTransactions() ([]*e.Transactions, error)
}
