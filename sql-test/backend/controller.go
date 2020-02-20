package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"net/http"
)

func checkdb(w http.ResponseWriter, r *http.Request) {
	// Test MSSQL connection
	pathParams := mux.Vars(r)
	ipAddress := pathParams["ipAddress"]
	username := pathParams["username"]
	password := pathParams["password"]
	port := 1433
	timeout := "2s"

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d; dial timeout=%s", ipAddress, username, password, port, timeout)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		response := Response{
			Message: "NOK",
			Status:  false,
		}
		json.NewEncoder(w).Encode(response)
	} else {
		response := Response{
			Message: "NOK",
			Status:  false,
		}
		json.NewEncoder(w).Encode(response)
	}
	defer conn.Close()
}
