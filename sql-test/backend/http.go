package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Create router
	r := mux.NewRouter()

	r.HandleFunc("/", OptionForCORS).Methods(http.MethodOptions)
	//r.HandleFunc("/", Homepage).Methods(http.MethodGet)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./"))).Methods(http.MethodGet)
	r.HandleFunc("/check", OptionForCORS).Methods(http.MethodOptions)
	r.HandleFunc("/check", CheckDB).Methods(http.MethodPost)

	// Start serving HTTP
	log.Fatal(http.ListenAndServe(":80", r))
}
