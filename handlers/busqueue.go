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

func busqueueData(routeid string) models.FullResp {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	conn, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	db := conn

	tsql := fmt.Sprintf(`
	SELECT BusesQueueID, BusOrderID,BusClassID, BusShapeID,BusSeatArrangementID,
	BusesQueueStatus
	FROM BusesQueue
	WHERE routeid = '%s'
	`, routeid)

	rows, err := db.Query(tsql)

	if err != nil {
		fmt.Println("error querying database")
	}

	defer rows.Close()

	var id, orderid, classid, shapeid, arrangementid, status string

	var tempRes []models.Resp
	var tempy models.Resp

	for rows.Next() {

		err := rows.Scan(&id, &orderid, &classid, &shapeid, &arrangementid, &status)

		if err != nil {
			fmt.Println("error looping rows")
		}

		tempy = models.Resp{
			"busqueueid":           id,
			"busorderid":           orderid,
			"busclassid":           classid,
			"busshapeid":           shapeid,
			"busseatarrangementid": arrangementid,
			"busqueuestatus":       status,
		}

		tempRes = append(tempRes, tempy)

	}

	var finRes models.FullResp
	finRes.Status = "found"
	finRes.Buses = tempRes

	return finRes
}

func Busqueue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	routeId := r.Header["Routeid"][0]

	data := busqueueData(routeId)

	json.NewEncoder(w).Encode(data)

	Logger("/v3/buses")
}
