package config

import (
	"database/sql"
	"log"
)

var db *sql.DB

func DatabaseInit() {
	var err error
	//TODO: set var by .env
	db, err = sql.Open("postgres", "user=test dbname=goapi password=test host=postgres port=5433 sslmode=disable")
	// db, err = sql.Open("postgres", "user=test dbname=goapi password=test host=localhost port=5432 sslmode=disable")

	if err != nil {
		log.Panic(err)
	}

	createToRewardBase()
}

func createToRewardBase() {

}
