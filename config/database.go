package config

import "database/sql"

var db *sql.DB

func DatabaseInit() {

	createToRewardBase()
}

func createToRewardBase() {

}
