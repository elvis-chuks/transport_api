package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	"./handlers"
)

var (
	server = "DESKTOP-NIT9OR4"
	port   = 1433
	user   = "sa"
	password = "@123elvischuks"
	database = "maadendum"
)

 

func test(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")

	res := map[string]string{"hello":"world"}
	json.NewEncoder(w).Encode(res)
}


func main(){
	fmt.Println("this is the pmt api written in go")
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server,user,password,port,database)

	conn,err := sql.Open("mssql",connString)

	if err != nil {
		log.Fatal("Open connection failed:",err.Error())
	}

	fmt.Printf("Connected!\n")
	
	defer conn.Close()
	

	router := mux.NewRouter()

	router.HandleFunc("/1",test).Methods("GET")
	router.HandleFunc("/v3/hello",handlers.Hello).Methods("GET")
	router.HandleFunc("/v3/zroute",handlers.Zroute).Methods("GET")
	

	log.Fatal(http.ListenAndServe(":5000", router))
}