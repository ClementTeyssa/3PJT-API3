package main

import (
	"log"
	"net/http"

	"github.com/ClementTeyssa/3PJT-API3/config"
)

func main() {
	config.DatabaseInit()
	log.Println("Database initialised")
	router := InitializeRouter()
	log.Println("Rooter initialised")
	log.Panic(http.ListenAndServe(":8080", router))
}
