package database

import (
	"database/sql"
)


// Denne må skrives om
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
