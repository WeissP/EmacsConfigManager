package connector

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	filenameFormat = "weiss_%s<%s.el"
)

func Connect() (db *sql.DB) {
	db, _ = sql.Open("sqlite3", "/home/weiss/weiss/EmacsConfigManager/connector/emacsconfig.db")
	return
}
