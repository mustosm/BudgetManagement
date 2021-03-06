package main

import (
	"log"
	"strconv"
	"net/http"
)

var config = Config{}
var dao = TransactionsDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.DBServer = config.DBServer
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	router := NewRouter()
	router.Use(RequestLoggerMiddleware)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), router))
	return
}
