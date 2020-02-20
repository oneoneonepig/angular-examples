package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"html/template"
	"log"
	"net/http"
)

type Request struct {
	IpAddress string `json:"inputIpAddress"`
	Username  string `json:"inputUsername"`
	Password  string `json:"inputPassword"`
	Timeout   int    `json:"inputTimeout"`
}

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func OptionForCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
	w.WriteHeader(http.StatusOK)
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	title := "SQL Connection Test"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, title)
}

func CheckDB(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "")
	// Decode the request body to json
	decoder := json.NewDecoder(r.Body)
	var input Request
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}
	port := 1433

	// Create connection to database
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d; connection timeout=%d; dial timeout=%d",
		input.IpAddress, input.Username, input.Password, port, input.Timeout, input.Timeout)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	// Ping the database,
	// response depends on the result of the ping
	err = conn.Ping()
	if err != nil {
		response := Response{
			Message: err.Error(),
			Status:  false,
		}
		json.NewEncoder(w).Encode(response)
	} else {
		response := Response{
			Message: "Connection OK",
			Status:  true,
		}
		json.NewEncoder(w).Encode(response)
	}
}
