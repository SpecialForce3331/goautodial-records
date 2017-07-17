package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"server/config_parser"
)

type Record struct {
    Agent string
    Phone string
    Location string
    CallDate string
    IsInbound bool
}

var cfg config_parser.Config

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
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",cfg.SQL_LOGIN, cfg.SQL_PASSWORD, cfg.SQL_HOST, cfg.SQL_PORT, cfg.SQL_DB))
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
			var listId string
			var isInbound bool = false

	        if err := rows.Scan(&agent, &phone, &location, &callDate, &listId); err != nil {
                log.Println(err)
                log.Println(agent, phone, location, callDate, listId)
	        }

	        // fmt.Println(listId, listId == "900")

	        if listId == cfg.SQL_FIELD_INBOUND_list_id {
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
	cfg = config_parser.GetConfig() 

	fmt.Println(fmt.Sprintf("Server started!\nListening %s port...",cfg.HTTP_PORT))
	http.HandleFunc("/records", handler)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(fmt.Sprintf(":%s",cfg.HTTP_PORT), nil)
}
