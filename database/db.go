package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

func ConnectDB() {
	loadEnv()
	DB, err = sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		panic("Error while connecting to the database.")
	}
	fmt.Println("Connected to the database!")
	createTables()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error while loading .env")
	}
}

func createTables() {
	createEventsTable := `
        CREATE TABLE IF NOT EXISTS events (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            dateTime TIMESTAMP NOT NULL,
            user_id INTEGER
        )
    `
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Error while creating events table.")
	} else {
		fmt.Println("Events table created.")
	}
}