package main

import (
	"database/sql"
	"fmt"
	"log"

	"encoding/json"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"

	"./handlers"
)

var (
	server   = "DESKTOP-NIT9OR4"
	port     = 1433
	user     = "sa"
	password = "@123elvischuks"
	database = "maadendum"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	res := map[string]string{"hello": "world"}
	json.NewEncoder(w).Encode(res)
}

func main() {
	fmt.Println("this is the pmt api written in go")

	// sql connection string

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	conn, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	fmt.Printf("Connected!\n")

	defer conn.Close()

	router := mux.NewRouter()

	router.HandleFunc("/1", test).Methods("GET")
	router.HandleFunc("/v3/hello", handlers.Hello).Methods("GET")
	router.HandleFunc("/v3/zroute", handlers.Zroute).Methods("GET")
	router.HandleFunc("/v3/discount", handlers.Discount).Methods("GET")
	router.HandleFunc("/v3/buses", handlers.Busqueue).Methods("GET")
	router.HandleFunc("/v3/seatamount", handlers.SeatAmnt).Methods("GET")
	router.HandleFunc("/v3/seats", handlers.Seats).Methods("GET")
	router.HandleFunc("/v3/trips", handlers.Trips).Methods("GET")
	router.HandleFunc("/v3/gettrips", handlers.GetTrips).Methods("GET")
	router.HandleFunc("/v3/checkbook", handlers.CheckBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
