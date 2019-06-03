package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DatabaseInit() {
	var err error
	//TODO: set var by .env
	db, err = sql.Open("postgres", "user=test dbname=goapi password=test host=postgres port=5432 sslmode=disable")
	// db, err = sql.Open("postgres", "user=test dbname=goapi password=test host=localhost port=5433 sslmode=disable")

	if err != nil {
		log.Panic(err)
	}

	createToRewardBase()
}

func createToRewardBase() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS torewards(id serial, adress varchar, number integer, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk_users primary key(id))")
	if err != nil {
		log.Panic(err)
	}
}
