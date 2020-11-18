package database

import (
	"database/sql"
)

func CheckInDatabase(website string, db *sql.DB) bool {
	statement, _ := db.Prepare("SELECT website FROM passwords")
	rows, _ := statement.Query()

	var websiteRes string

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&websiteRes)

		if websiteRes == website {
			return true
		}
	}
	return false
}
