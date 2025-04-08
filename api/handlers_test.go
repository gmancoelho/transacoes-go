package api

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestHandleTransaction(t *testing.T) {
    server := &APIServer{}

    t.Run("Given a valid POST request to /transaction", func(t *testing.T) {
        // Arrange
        req := httptest.NewRequest(http.MethodPost, "/transaction", strings.NewReader(`{"value": 100, "dateHour": "2025-04-08T15:04:05Z"}`))
        w := httptest.NewRecorder()

        // Act
        err := server.handleTransaction(w, req)

        // Assert
        assert.NoError(t, err, "unexpected error occurred")
        assert.Equal(t, http.StatusOK, w.Code, "expected status code 200")
    })

    t.Run("Given a valid DELETE request to /transaction", func(t *testing.T) {
        // Arrange
        req := httptest.NewRequest(http.MethodDelete, "/transaction", nil)
        w := httptest.NewRecorder()

        // Act
        err := server.handleTransaction(w, req)

        // Assert
        assert.NoError(t, err, "unexpected error occurred")
        assert.Equal(t, http.StatusOK, w.Code, "expected status code 200")
    })

    t.Run("Given an invalid method PUT for /transaction", func(t *testing.T) {
        // Arrange
        req := httptest.NewRequest(http.MethodPut, "/transaction", nil)
        w := httptest.NewRecorder()

        // Act
        err := server.handleTransaction(w, req)

        // Assert
        assert.Error(t, err, "expected an error but got none")
        assert.Equal(t, http.StatusMethodNotAllowed, w.Code, "expected status code 405")
    })
}

func TestHandleStatistic(t *testing.T) {
    server := &APIServer{}

    t.Run("Given a valid GET request to /statistic", func(t *testing.T) {
        // Arrange
        req := httptest.NewRequest(http.MethodGet, "/statistic", nil)
        w := httptest.NewRecorder()

        // Act
        err := server.handleStatistic(w, req)

        // Assert
        assert.NoError(t, err, "unexpected error occurred")
        assert.Equal(t, http.StatusOK, w.Code, "expected status code 200")
    })

    t.Run("Given an invalid method POST for /statistic", func(t *testing.T) {
        // Arrange
        req := httptest.NewRequest(http.MethodPost, "/statistic", nil)
        w := httptest.NewRecorder()

        // Act
        err := server.handleStatistic(w, req)

        // Assert
        assert.Error(t, err, "expected an error but got none")
        assert.Equal(t, http.StatusMethodNotAllowed, w.Code, "expected status code 405")
    })
}