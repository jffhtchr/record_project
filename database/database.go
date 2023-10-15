package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Dbinstance struct {
	Db *sql.DB
}

var DB Dbinstance

func ConnectDatabase() {
	// Connect to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS records (id SERIAL PRIMARY KEY, title TEXT, artist TEXT, price FLOAT, available BOOL)")
	if err != nil {
		log.Fatal(err)
	}

	DB = Dbinstance{
		Db: db,
	}
}
