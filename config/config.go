package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetMySQLDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "todo"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)
	return
}
