package api

import (
	"encoding/json"
	"log"
	"net/http"

	m "github.com/gmancoelho/transacoes-go/entities"
	s "github.com/gmancoelho/transacoes-go/repository"
	"github.com/gorilla/mux"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type APIServer struct {
	address string
	storage s.Storage
}

func writeJSON(w http.ResponseWriter, statusCode int, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(response)
}

func makeHTTPHandlerFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError,
				m.ApiError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
		}
	}
}

func NewAPIServer(add string, storage s.Storage) *APIServer {
	return &APIServer{
		address: add,
		storage: storage,
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (s *APIServer) Start() error {
	router := s.setupRouter()

	log.Printf("API server started on %s\n", s.address)

	return http.ListenAndServe(s.address, router)
}

func (s *APIServer) setupRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(LoggingMiddleware)

	router.HandleFunc("/transacao", makeHTTPHandlerFunc(s.handleTransaction))
	router.HandleFunc("/estatistica", makeHTTPHandlerFunc(s.handleTransaction))

	return router
}
