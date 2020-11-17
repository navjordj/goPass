package database

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "reflect"
	
)
/*
type database struct {
	s string
}*/

func NewDatabase(name string) { //**sql.DB {
	//db := database{s: name}

	db, _ := sql.Open("sqlite3", name)

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS passwords (id INTEGER PRIMARY KEY, password Text, website Text)")
	statement.Exec()

	//return &db
}

func Insert(password string, website string, db_name string) int {

	db, _ := sql.Open("sqlite3", db_name)

	res := CheckInDatabase(website, db)

	if (res == true) {
		fmt.Println("\nWebsite already in DB")
		return -1
	} else {
		statement2, _ := db.Prepare("INSERT INTO passwords (password, website) VALUES (?, ?)")
		statement2.Exec(password, website)
		return 1
	}
	
}

func Get(website string, db *sql.DB) string {

	rows, _ := db.Query("SELECT password from passwords WHERE website = ?", website)

	var password string
	for rows.Next() {
		rows.Scan(&password)
	}
	return password
}