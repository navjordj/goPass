package database

import (
	_"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
/*
type database struct {
	s string
}*/

func NewDatabase(name string) **sql.DB {
	//db := database{s: name}

	db, _ := sql.Open("sqlite3", name)

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS passwords (id INTEGER PRIMARY KEY, password Text, website Text)")
	statement.Exec()

	return &db
}

func Insert(password string, website string, db *sql.DB) int {
	statement, _ := db.Prepare("INSERT INTO passwords (password, website) VALUES (?, ?)")
	statement.Exec(password, website)
	return 1
}