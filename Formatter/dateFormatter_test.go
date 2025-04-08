package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertAValidISO8601ToString(t *testing.T) {
	t.Run("Given a valid ISO 8601 date", func(t *testing.T) {
		// Arrange
		isoDate := "2025-04-08T15:04:05Z"
		expected := "08/04/2025 15:04:05"

		// Act
		result, err := ConvertISO8601ToString(isoDate)

		// Assert
		assert.NoError(t, err, "unexpected error occurred")
		assert.Equal(t, expected, result, "formatted date does not match expected value")
	})
}
func TestConvertAnInValidISO8601ToString(t *testing.T) {
	t.Run("Given an invalid ISO 8601 date", func(t *testing.T) {
		// Arrange
		isoDate := "invalid-date"

		// Act
		result, err := ConvertISO8601ToString(isoDate)

		// Assert
		assert.Error(t, err, "expected an error but got none")
		assert.Empty(t, result, "expected an empty string but got a value")
	})
}

func TestConvertAnEmptyISO8601ToString(t *testing.T) {
	t.Run("Given an empty ISO 8601 date", func(t *testing.T) {
		// Arrange
		isoDate := ""

		// Act
		result, err := ConvertISO8601ToString(isoDate)

		// Assert
		assert.Error(t, err, "expected an error but got none")
		assert.Empty(t, result, "expected an empty string but got a value")
	})
}
