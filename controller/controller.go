package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"postgres/models"
)

// func GetUsers(db *sql.DB) []models.User {
// 	rows, err := db.Query("SELECT id, username FROM users")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var users []models.User
// 	for rows.Next() {
// 		var u models.User
// 		if err := rows.Scan(&u.ID, &u.Username); err != nil {
// 			log.Fatal(err)
// 		}
// 		users = append(users, u)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

//		return users
//	}
var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Use the db to query the database...
	rows, err := db.Query("SELECT id, username FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
