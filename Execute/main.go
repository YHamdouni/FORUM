package main

import (
	"database/sql"
	"log"
	Files "main/Handlers/Files"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS User (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		deja INT DEFAULT 0
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
	InitDB()
	Files.RegisterRoutes(db)
	// fmt.Println("Server running at http://localhost:8080")
	// err = http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	port := os.Getenv("PORT")
   	if port == "" {
        port = "8080" // Default to port 8080 if not set
        	log.Println("No PORT environment variable detected, defaulting to", port)
    	}
	log.Println("Starting server on port", port)
    	if err := http.ListenAndServe(":" + port, nil); err != nil {
        	log.Fatal(err)
    	}
}
func InitDB() {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS Posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
		category TEXT NOT NULL,
        title TEXT UNIQUE NOT NULL,
        content TEXT NOT NULL,
		date TEXT, 
		time TIME
    );
    `)
	if err != nil {
		log.Fatal("Error creating Posts table: ", err)
	}
}
