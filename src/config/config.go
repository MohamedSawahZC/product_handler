package config

import "database/sql"

func GetDB(db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "dev"
	db, err := sql.Open(dbDriver, dbUser, ":", dbPass, "@/", dbName)
	return
}
