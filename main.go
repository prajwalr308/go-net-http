package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"postgres/controller"
)

func main() {
	err := godotenv.Load()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	fmt.Printf("Connecting with: host=%s user=%s dbname=%s sslmode=disable\n", dbHost, dbUser, dbName)
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPass, dbName)
	db, err := sql.Open("postgres", connStr)

	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(username)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	controller.SetDB(db)
	http.HandleFunc("/users", controller.GetUsers)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
