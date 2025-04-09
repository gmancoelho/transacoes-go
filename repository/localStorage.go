package storage

import e "github.com/gmancoelho/transacoes-go/entities"

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
	return nil
}

func (s *LocalStorage) DeleteTransaction(id int) error {
	return nil
}

func (s *LocalStorage) GetTransactions() ([]*e.Transactions, error) {
	return nil, nil
}
