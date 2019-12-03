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



func zrouteData(r string) models.FullResp {
	
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server,user,password,port,database)

	conn,err := sql.Open("mssql",connString)

	if err != nil {
		log.Fatal("Open connection failed:",err.Error())
	}

	db := conn
	
	tsql := fmt.Sprintf( `
	SELECT zrouteid,dru
	FROM dbo.zroute
	WHERE oru = '%s';
	`,r)

	rows,err := db.Query(tsql)

	if err != nil{
		fmt.Println("Error reading rows")
	}

	defer rows.Close()
	var res []models.Resp
	var tempRes []models.Resp
	for rows.Next() {
		var id,dru string
		err := rows.Scan(&id,&dru)

		if err != nil{
			fmt.Println("Error looping rows")
		}
		payload := models.Resp{"id":id,"dru":dru}
		tempRes = append(tempRes,payload)
	}

	for _,i := range tempRes {
		tsql = fmt.Sprintf(`
			SELECT depotname
			FROM dbo.depot
			WHERE depotcode = '%s';
			`,i["dru"])
		
			rows,err := db.Query(tsql)

			if err != nil{
				fmt.Println("Error reading rows")
			}
		
			defer rows.Close()

			

			for rows.Next() {
				var name string
				err := rows.Scan(&name)
				if err != nil {
					fmt.Println("error looping through rows")
				}

				resp := models.Resp{"routeid":i["id"],"routename":name}
				res = append(res,resp)

			}


	}

	var finRes models.FullResp

	finRes.Status = "found"
	finRes.Routes = res

	return finRes

}

func Zroute(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")

	

	depotcode := r.Header["Depotcode"][0]
	
	// check if depotcode is empty

	data := zrouteData(depotcode)

	json.NewEncoder(w).Encode(data)

	Logger("/v3/zroute")
}