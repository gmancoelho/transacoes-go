package api

import (
	"fmt"
	"net/http"

	m "github.com/gmancoelho/transacoes-go/models"
)

func (s *APIServer) handleTransaction(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodPost:
		return s.handleCreateTrasaction(w, r)
	case http.MethodDelete:
		return s.handleDeleteAccount(w, r)
	default:
		return handleError(w, fmt.Errorf("method not allowed: %s", r.Method), http.StatusMethodNotAllowed)
	}
}

func (s *APIServer) handleStatistic(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetStatistics(w, r)
	default:
		return handleError(w, fmt.Errorf("method not allowed: %s", r.Method), http.StatusMethodNotAllowed)
	}
}

func (s *APIServer) handleGetStatistics(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateTrasaction(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func handleError(w http.ResponseWriter, err error, statusCode int) error {
	writeJSON(w, statusCode, m.ApiError{
		Code:    statusCode,
		Message: err.Error(),
	})
	return err
}
