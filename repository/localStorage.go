package repository

import (
	"fmt"

	e "github.com/gmancoelho/transacoes-go/entities"
)

type LocalStorage struct {
	transactions []e.Transactions
}

func NewLocalStorage() *LocalStorage {
	transactionsSlice := make([]e.Transactions, 0)

	return &LocalStorage{
		transactions: transactionsSlice,
	}
}

func (s *LocalStorage) CreateTransaction(transaction e.Transactions) error {
	s.transactions = append(s.transactions, transaction) // Update the slice in place
	return nil
}

func (s *LocalStorage) DeleteTransaction(id int) error {
	if len(s.transactions) == 0 {
		return fmt.Errorf("transactions is empty")
	}

	if id < 0 || id >= len(s.transactions) {
		return fmt.Errorf("transaction id not found")
	}

	s.transactions = append(s.transactions[:id], s.transactions[id+1:]...)

	return nil
}

func (s *LocalStorage) GetTransactions() []e.Transactions {
	return s.transactions
}
