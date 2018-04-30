package main

import (
	"log"
	"net/http"
)

var config = Config{}
var dao = TransactionsDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	router := NewRouter()
	router.Use(RequestLoggerMiddleware)
	log.Fatal(http.ListenAndServe(":8080", router))

	return
}
