package config_parser

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const configName string = "config.txt"

type Config struct {
	HTTP_PORT                 string
	SQL_LOGIN                 string
	SQL_PASSWORD              string
	SQL_HOST                  string
	SQL_PORT                  string
	SQL_DB                    string
	SQL_FIELD_INBOUND_list_id string
	DAYS_SELECT_COUNT         string
}

func openFile() *os.File {
	file, err := os.Open(configName)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func parseConfig(file *os.File) Config {

	var HTTP_PORT string
	var SQL_LOGIN string
	var SQL_PASSWORD string
	var SQL_HOST string
	var SQL_PORT string
	var SQL_DB string
	var SQL_FIELD_INBOUND_list_id string
	var DAYS_SELECT_COUNT string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineArguments := strings.Split(line, "=")

		switch key := lineArguments[0]; key {
		case "HTTP_PORT":
			HTTP_PORT = lineArguments[1]
		case "SQL_LOGIN":
			SQL_LOGIN = lineArguments[1]
		case "SQL_PASSWORD":
			SQL_PASSWORD = lineArguments[1]
		case "SQL_HOST":
			SQL_HOST = lineArguments[1]
		case "SQL_PORT":
			SQL_PORT = lineArguments[1]
		case "SQL_DB":
			SQL_DB = lineArguments[1]
		case "SQL_FIELD_INBOUND_list_id":
			SQL_FIELD_INBOUND_list_id = lineArguments[1]
		case "DAYS_SELECT_COUNT":
			DAYS_SELECT_COUNT = lineArguments[1]
		default:
			log.Fatal("Config error! Please check config.txt file!")
		}
	}

	if HTTP_PORT == "" || SQL_LOGIN == "" || SQL_PASSWORD == "" || SQL_HOST == "" || SQL_PORT == "" || SQL_DB == "" || SQL_LOGIN == "" || SQL_FIELD_INBOUND_list_id == "" || DAYS_SELECT_COUNT == "" {
		log.Fatal("Config error! One of fields are null")
	}

	return Config{HTTP_PORT, SQL_LOGIN, SQL_PASSWORD, SQL_HOST, SQL_PORT, SQL_DB, SQL_FIELD_INBOUND_list_id, DAYS_SELECT_COUNT}

}

func GetConfig() Config {
	file := openFile()
	defer file.Close()

	return parseConfig(file)
}
