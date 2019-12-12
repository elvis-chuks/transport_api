package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"
	_ "github.com/denisenkom/go-mssqldb"
)

func tripData(id string, classId string) models.FullResp {
	// recieves routeid

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	conn, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	db := conn

	tsql := fmt.Sprintf(`
	SELECT BusFareCode,Prizing,BusFareStatus
	FROM dbo.BusFare
	WHERE RouteID = '%s' AND BusClassID = '%s'
	`, id, classId)

	rows, err := db.Query(tsql)

	if err != nil {
		fmt.Println("error querying database")
	}

	defer rows.Close()

	var BusFareCode, Prizing, BusFareStatus string
	var tempRes []models.Resp
	var tempy models.Resp
	for rows.Next() {

		err := rows.Scan(&BusFareCode, &Prizing, &BusFareStatus)

		if err != nil {
			fmt.Println("error looping rows")
		}

		tempy = models.Resp{
			"BusFareCode":   BusFareCode,
			"Prizing":       Prizing,
			"BusFareStatus": BusFareStatus,
		}

		tempRes = append(tempRes, tempy)

	}

	var finRes models.FullResp
	finRes.Status = "found"
	finRes.Prices = tempRes
	return finRes
}

func Trips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	routeId := r.Header["Routeid"][0]
	classId := r.Header["Busclassid"][0]

	data := tripData(routeId, classId)

	json.NewEncoder(w).Encode(data)

	Logger("/v3/trips")
}
