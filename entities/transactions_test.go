package entities

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionEntity(t *testing.T) {
	t.Run("Given a valid JSON from a request to /transaction", func(t *testing.T) {
		// Arrange
		req := strings.NewReader(`{"valor": 100, "dataHora": "2025-04-08T15:04:05Z"}`)

		// Act
		var tranRequest Transactions

		if err := json.NewDecoder(req).Decode(&tranRequest); err != nil { // Pass a pointer to tranRequest
			t.Fatalf("failed to decode request: %v", err)
		}

		// Assert
		assert.Equal(t, "2025-04-08T15:04:05Z", tranRequest.DateHour, "DateHour should match the input JSON")
		assert.Equal(t, 100, tranRequest.Value, "Value should match the input JSON") // Added assertion for Value
	})

	t.Run("Given a invalid JSON from a request to /transaction", func(t *testing.T) {
		// Arrange
		req := strings.NewReader(`{}`)

		// Act
		var tranRequest Transactions

		if err := json.NewDecoder(req).Decode(&tranRequest); err != nil { // Pass a pointer to tranRequest
			t.Fatalf("failed to decode request: %v", err)
		}

		// Assert
		assert.Equal(t, "", tranRequest.DateHour, "DateHour should match the input JSON")
		assert.Equal(t, 0, tranRequest.Value, "Value should match the input JSON") // Added assertion for Value
	})
}
