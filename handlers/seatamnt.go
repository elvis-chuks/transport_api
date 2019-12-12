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

func seatData(id string) models.FullResp {
	// recieves busseatarrangementid

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	conn, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	db := conn

	tsql := fmt.Sprintf(`
	SELECT NumberOfSeat,SeatMap,SeatCardinal
	FROM dbo.BusSeatArrangement
	WHERE BusSeatArrangementID = '%s'
	`, id)

	rows, err := db.Query(tsql)

	if err != nil {
		fmt.Println("error querying database")
	}

	defer rows.Close()

	var noOfSeat, seatMap, seatCardinal string
	var tempRes []models.Resp
	var tempy models.Resp
	for rows.Next() {

		err := rows.Scan(&noOfSeat, &seatMap, &seatCardinal)

		if err != nil {
			fmt.Println("error looping rows")
		}

		tempy = models.Resp{
			"NumberOfSeat": noOfSeat,
			"SeatMap":      seatMap,
			"SeatCardinal": seatCardinal,
		}

		tempRes = append(tempRes, tempy)

	}

	var finRes models.FullResp
	finRes.Status = "found"
	finRes.BusSeat = tempRes
	return finRes
}

func SeatAmnt(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	seatId := r.Header["Busseatarrangementid"][0]

	data := seatData(seatId)

	json.NewEncoder(w).Encode(data)

	Logger("/v3/seatamount")
}
