package repository

import (
	"testing"

	e "github.com/gmancoelho/transacoes-go/entities"
	"github.com/stretchr/testify/assert"
)

func TestLocalStorage(t *testing.T) {
	t.Run("Given a call to NewLocalStorage it should return a valid Storage", func(t *testing.T) {
		// Act
		storage := NewLocalStorage()

		// Assert
		assert.NotNil(t, storage, "NewLocalStorage should return a non-nil instance")
		assert.Empty(t, storage.transactions, "NewLocalStorage should initialize with an empty transactions slice")
	})

	t.Run("Given a valid transaction, CreateTransaction should add it to storage", func(t *testing.T) {
		// Arrange
		storage := NewLocalStorage()
		transaction := e.Transactions{
			Value:    100,
			DateHour: "2025-04-08T15:04:05Z",
		}

		// Act
		err := storage.CreateTransaction(transaction)

		// Assert
		assert.NoError(t, err, "CreateTransaction should not return an error")
		assert.Len(t, storage.transactions, 1, "Storage should contain one transaction after creation")
		assert.Equal(t, transaction, storage.transactions[0], "Stored transaction should match the input transaction")
	})

	t.Run("Given a transaction ID, DeleteTransaction should remove it from storage", func(t *testing.T) {
		// Arrange
		storage := NewLocalStorage()
		transaction := e.Transactions{
			Value:    100,
			DateHour: "2025-04-08T15:04:05Z",
		}
		storage.transactions = append(storage.transactions, transaction)

		// Act
		err := storage.DeleteTransaction(0)

		// Assert
		assert.NoError(t, err, "DeleteTransaction should not return an error")
		assert.Empty(t, storage.transactions, "Storage should be empty after deleting the transaction")
	})

	t.Run("GetTransactions should return all stored transactions", func(t *testing.T) {
		// Arrange
		storage := NewLocalStorage()
		transaction1 := e.Transactions{
			Value:    100,
			DateHour: "2025-04-08T15:04:05Z",
		}
		transaction2 := e.Transactions{
			Value:    200,
			DateHour: "2025-04-09T15:04:05Z",
		}
		storage.transactions = append(storage.transactions, transaction1, transaction2)

		// Act
		transactions := storage.GetTransactions()

		// Assert
		assert.Len(t, transactions, 2, "GetTransactions should return all stored transactions")
		assert.Equal(t, transaction1, transactions[0], "First transaction should match the stored transaction")
		assert.Equal(t, transaction2, transactions[1], "Second transaction should match the stored transaction")
	})
}
