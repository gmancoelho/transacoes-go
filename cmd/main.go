package main

import (
	"fmt"

	"github.com/gmancoelho/transacoes-go/api"
	s "github.com/gmancoelho/transacoes-go/repository"
)

const address = ":8080"

func main() {
	fmt.Printf("Start Transactions Server \n")

	store := s.NewLocalStorage()

	server := api.NewAPIServer(address, store)

	if err := server.Start(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

}
