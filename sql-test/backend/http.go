package backend

import (
	"github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	connString = ""
	// Example: "sqlserver://username:password@host/instance?param1=value&param2=value"
)

func main() {
	// Create router
	r := mux.NewRouter()

	r.HandleFunc("/checkdb", test).Methods(http.MethodPost)

	// Create redis connection
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

	// Test redis connection
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}

	// Start serving HTTP
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
