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

	statement, _ := db.Prepare("SELECT website FROM passwords")
	rows, _ := statement.Query()
	hva, _ := rows.ColumnTypes()
	fmt.Println(hva)

	var website_res string

	for rows.Next() {
		rows.Scan(&website_res)

		if (website_res == website) {
			return -1
		} 

	}


	statement2, _ := db.Prepare("INSERT INTO passwords (password, website) VALUES (?, ?)")
	statement2.Exec(password, website)
	return 1
}