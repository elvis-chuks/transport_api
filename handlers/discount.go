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

func discountData(routeid string) models.Discount {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	conn, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	db := conn

	tsql := fmt.Sprintf(`
	SELECT busclassid,busorderid,discountid
	FROM dbo.routediscount
	WHERE routeid = '%s'
	`, routeid)

	rows, err := db.Query(tsql)

	if err != nil {
		fmt.Println("error querying database")
	}

	defer rows.Close()

	var busclassid, busorderid, discountid string

	for rows.Next() {

		err := rows.Scan(&busclassid, &busorderid, &discountid)

		if err != nil {
			fmt.Println("error looping rows")
		}

	}
	// fmt.Println(discountid, "discount id")
	tsql = fmt.Sprintf(`
	SELECT discountname,discountamount
	FROM dbo.discount
	WHERE discountid = '%s'
	`, discountid)
	// fmt.Println(tsql)
	rows, err = db.Query(tsql)

	if err != nil {
		fmt.Println("error querying database")
	}

	defer rows.Close()

	var discountname, discountamount string

	for rows.Next() {
		err := rows.Scan(&discountname, &discountamount)

		if err != nil {
			fmt.Println("error looping rows")
		}

	}
	// create return payload

	var finres models.Discount

	finres.Status = "found"
	finres.Busclassid = busclassid
	finres.Busorderid = busorderid
	finres.Discountid = discountid
	finres.Discountname = discountname
	finres.Discountamount = discountamount

	return finres
}

func Discount(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	routeId := r.Header["Routeid"][0]

	data := discountData(routeId)

	json.NewEncoder(w).Encode(data)

	Logger("/v3/discount")
}
