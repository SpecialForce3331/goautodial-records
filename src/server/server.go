package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

const (
	HTTP_PORT int = 8080
	SQL_LOGIN string = "records"
	SQL_PASSWORD string = "283g238dg28g"
	SQL_HOST string = "localhost"
	SQL_PORT int = 3306
	SQL_DB string = "asterisk"
	SQL_FIELD_INBOUND_list_id int = 999
) 

type Record struct {
    Agent string
    Phone string
    Location string
    CallDate string
    IsInbound bool
}

func handler(w http.ResponseWriter, r *http.Request) {
	records := getRecords()
	records_json, err := json.Marshal(records)

	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
  	w.Write(records_json)
}

func dbOpen() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",SQL_LOGIN, SQL_PASSWORD, SQL_HOST, SQL_PORT, SQL_DB))
	if err != nil {
    	panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}

func dbSelect(db *sql.DB) []Record {
	rows, err := db.Query("SELECT agent, phone, location, call_date, list_id FROM goautodial_recordings_views WHERE location IS NOT NULL AND phone IS NOT NULL and list_id IS NOT NULL ORDER BY call_date DESC")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer rows.Close()

	var records []Record

	for rows.Next() {
	        var agent string
			var phone string
			var location string
			var callDate string
			var listId int
			var isInbound bool = false


	        if err := rows.Scan(&agent, &phone, &location, &callDate, &listId); err != nil {
                log.Println(err)
                log.Println(agent, phone, location, callDate, listId)
	        }

	        if listId == SQL_FIELD_INBOUND_list_id  {
	        	isInbound = true
	        }

	        record := Record{agent, phone, location, callDate, isInbound}

	        records = append(records, record)

	        
	}

	if err := rows.Err(); err != nil {
	        log.Fatal(err)
	}

	// fmt.Println(records)
	return records
}

func getRecords() []Record {
	db := dbOpen()
	defer db.Close()

	return dbSelect(db)
}

func main() {
	fmt.Println(fmt.Sprintf("Server started!\nListening %d port...",HTTP_PORT))
	http.HandleFunc("/records", handler)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(fmt.Sprintf(":%d",HTTP_PORT), nil)
}
