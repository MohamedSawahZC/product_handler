package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	dbUser := "root"
	dbPass := "root"
	dbName := "dev"
	dbQuery := dbUser + ":" + dbPass + "@/" + dbName
	fmt.Println(dbQuery)
	db, err = sql.Open("mysql", "root:root@/dev")
	return
}
