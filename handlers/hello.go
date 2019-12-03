package handlers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"encoding/json"
	"net/http"
	"../models"

)

var (
	server = "DESKTOP-NIT9OR4"
	port   = 1433
	user   = "sa"
	password = "@123elvischuks"
	database = "maadendum"
)

func helloData() models.FullResp{

	// sql connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server,user,password,port,database)

	conn,err := sql.Open("mssql",connString)

	if err != nil {
		log.Fatal("Open connection failed:",err.Error())
	}

	db := conn
	// sql query
	tsql := `
	SELECT oru FROM dbo.zroute
	GROUP BY oru;
	`
	rows,err := db.Query(tsql)

	if err != nil{
		fmt.Println("Error reading rows")
	}

	defer rows.Close()

	// create empty list
	var oruList []string
	// iterate through rows and drop the values into the oruList array
	for rows.Next(){
		var oru string

		err := rows.Scan(&oru)

		if err != nil{
			fmt.Println("Error looping rows")
		}

		oruList = append(oruList,oru)

	}

	var list []models.Resp
	
	// iterate through the oruList array and query the database for depot details
	for _,i := range oruList {

		tsql = fmt.Sprintf(`
	SELECT depotid, depotname,depotnumber,bankaccountnumber FROM dbo.Depot
	WHERE depotcode = '%s'
	ORDER BY DepotName ASC;
	`,i)

	rows,err := db.Query(tsql)

	if err != nil {
		fmt.Println("Error looping rows")
	}
	defer rows.Close()

	for rows.Next() {
		var id,name,number,account string

		err := rows.Scan(&id,&name,&number,&account)

		if err != nil {
			fmt.Println("error Looping rows")
		}

		// create a map of type models.Resp and append it to

		res := models.Resp{"depotid":id,"name":name,"depotcode":i,"number":number,"bankaccountnumber":account}

		list = append(list,res)
	}

	}



	var FinRes models.FullResp
	FinRes.Status = "found"
	FinRes.Depot = list

	return FinRes
	
} 


func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","application/json")

	
	
	data := helloData()

	json.NewEncoder(w).Encode(data)

	Logger("/v3/home") // call the logger helper function
}