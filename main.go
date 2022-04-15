package main

import (
	"os"

	"github.com/wdoogz/myRetail-RESTful-API/api"
	"github.com/wdoogz/myRetail-RESTful-API/db_connector"
)

func main() {
	_, exists := os.LookupEnv("LOADDB")
	if exists {
		db_connector.LoadDB()
	}
	api.Handle("9999")
}
