package entities

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type transactionId struct {
	Id int `json:"id"`
}

func ParseTransactionId(r *http.Request) (transactionId, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return transactionId{}, fmt.Errorf("invalid or missing account ID")
	}
	transactionId := transactionId{Id: id}
	return transactionId, nil
}
