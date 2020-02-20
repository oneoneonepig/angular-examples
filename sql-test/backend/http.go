package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Create router
	r := mux.NewRouter()

	r.HandleFunc("/checkdb", checkdb).Methods(http.MethodGet)

	// Start serving HTTP
	log.Fatal(http.ListenAndServe(":8888", r))
}

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
