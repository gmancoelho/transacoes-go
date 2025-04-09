package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	e "github.com/gmancoelho/transacoes-go/entities"
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
	defer r.Body.Close()

	tranRequest := new(e.Transactions)

	if err := json.NewDecoder(r.Body).Decode(tranRequest); err != nil {
		return handleError(w, fmt.Errorf("invalid request payload"), http.StatusBadRequest)
	}

	if tranRequest.DateHour == "" {
		return handleError(w, fmt.Errorf("date cannot be nil"), http.StatusBadRequest)
	}

	account := e.NewTransaction(tranRequest.Value, tranRequest.DateHour)

	return writeJSON(w, http.StatusCreated, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	_, err := e.ParseTransactionId(r)

	if err != nil {
		return handleError(w, fmt.Errorf("invalid request payload"), http.StatusBadRequest)
	}

	return writeJSON(w, http.StatusNoContent, nil)
}

func handleError(w http.ResponseWriter, err error, statusCode int) error {
	writeJSON(w, statusCode, e.ApiError{
		Code:    statusCode,
		Message: err.Error(),
	})
	return err
}
