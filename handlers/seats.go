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

func seatsData(id string) models.FullResp {
	// recieves busqueueid

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	conn, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	db := conn

	tsql := fmt.Sprintf(`
	SELECT BookedSeatID,TripDate,SeatNumber,SeatCardinal,BlockedStatus
	FROM dbo.BookedSeat
	WHERE BusesQueueID = '%s'
	`, id)

	rows, err := db.Query(tsql)

	if err != nil {
		fmt.Println("error querying database")
	}

	defer rows.Close()

	var seatid, tripdate, seatnumber, seatCardinal, blockedStatus string
	var tempRes []models.Resp
	var tempy models.Resp
	for rows.Next() {

		err := rows.Scan(&seatid, &tripdate, &seatnumber, &seatCardinal, &blockedStatus)

		if err != nil {
			fmt.Println("error looping rows")
		}

		tempy = models.Resp{
			"BookedSeatID":  seatid,
			"TripDate":      tripdate,
			"SeatNumber":    seatnumber,
			"SeatCardinal":  seatCardinal,
			"BlockedStatus": blockedStatus,
		}

		tempRes = append(tempRes, tempy)

	}

	var finRes models.FullResp
	finRes.Status = "found"
	finRes.Seats = tempRes
	return finRes
}

func Seats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	busId := r.Header["Busqueueid"][0]

	data := seatsData(busId)

	json.NewEncoder(w).Encode(data)

	Logger("/v3/seats")
}
