package entities

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestParseTransactionId(t *testing.T) {
	t.Run("Given a valid transaction ID in the request URL", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/transaction/123", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "123"})

		// Act
		transactionId, err := ParseTransactionId(req)

		// Assert
		assert.NoError(t, err, "parseTransactionId should not return an error for a valid ID")
		assert.Equal(t, 123, transactionId.Id, "Transaction ID should match the ID in the URL")
	})

	t.Run("Given an invalid transaction ID in the request URL", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/transaction/abc", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "abc"})

		// Act
		transactionId, err := ParseTransactionId(req)

		// Assert
		assert.Error(t, err, "parseTransactionId should return an error for an invalid ID")
		assert.Equal(t, 0, transactionId.Id, "Transaction ID should be 0 for an invalid ID")
	})

	t.Run("Given a missing transaction ID in the request URL", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/transaction/", nil)
		req = mux.SetURLVars(req, map[string]string{})

		// Act
		transactionId, err := ParseTransactionId(req)

		// Assert
		assert.Error(t, err, "parseTransactionId should return an error for a missing ID")
		assert.Equal(t, 0, transactionId.Id, "Transaction ID should be 0 for a missing ID")
	})
}
