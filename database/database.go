package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "reflect"

	_ "github.com/mattn/go-sqlite3"
)

/*
type database struct {
	s string
}*/

func NewDatabase(name string) { //**sql.DB {
	//db := database{s: name}

	db, _ := sql.Open("sqlite3", name)

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS passwords (id INTEGER PRIMARY KEY, password Text, email Text, website Text)")
	statement.Exec()

	//return &db
}

func Insert(password string, email string, website string, db_name string) int {

	db, _ := sql.Open("sqlite3", db_name)

	res := CheckInDatabase(website, db)

	if res == true {
		fmt.Println("\nWebsite already in DB")
		return -1
	} else {
		statement2, _ := db.Prepare("INSERT INTO passwords (password, email, website) VALUES (?, ?, ?)")
		statement2.Exec(password, email, website)
		return 1
	}
}

func Get(website string, db *sql.DB) (string, string) {

	rows, _ := db.Query("SELECT password, email from passwords WHERE website = ?", website)
	var password string
	var email string

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&password, &email)
	}

	return password, email
}

func Update(website string, newPassword string, db_name string) bool {

	db, _ := sql.Open("sqlite3", db_name)

	inDatabase := CheckInDatabase(website, db)

	if inDatabase {
		_, err := db.Exec("UPDATE passwords SET password = (?) WHERE website = (?)", newPassword, website)

		if err != nil {
			log.Fatal(err)
		}

		return true
	} else {
		fmt.Println("Ikke i databasen")
		return false
	}
}
