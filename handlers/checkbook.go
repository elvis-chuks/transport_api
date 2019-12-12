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

func checkBookData(phonenumber string) models.Resp {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	conn, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	db := conn

	tsql := fmt.Sprintf(`
	SELECT title,firstname,surname,gender,
	dateofbirth,emailaddress,nextofkinfullname,
	nextofkinphonenumber
	FROM dbo.CustomerDetails
	WHERE PhoneNumber = '%s'
	`, phonenumber)

	rows, err := db.Query(tsql)

	if err != nil {
		fmt.Println("error querying database")
	}

	defer rows.Close()

	var title, firstname, lastname, gender, dob, email, nextname, nextnumber string

	for rows.Next() {

		err := rows.Scan(&title, &firstname, &lastname, &gender, &dob, &email, &nextname, &nextnumber)

		if err != nil {
			fmt.Println("error looping rows")
		}

	}
	if len(title) <= 0 {
		return models.Resp{
			"status": "not found",
		}
	}
	return models.Resp{
		"status":               "found",
		"title":                title,
		"firstName":            firstname,
		"lastName":             lastname,
		"gender":               gender,
		"DOB":                  dob,
		"phoneNumber":          phonenumber,
		"email":                email,
		"nextOfKinFullName":    nextname,
		"nextOfKinPhoneNumber": nextnumber,
	}
}

func CheckBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	number := r.Header["Phonenumber"][0]

	data := checkBookData(number)

	json.NewEncoder(w).Encode(data)

	Logger("/v3/checkBook")
}
