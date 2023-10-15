package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	album "github.com/jffhtchr/record_project/pkg/models"
)

func GetAllAlbums(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Getting All Albums...")

		rows, err := db.Query("SELECT * FROM records")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		records := []album.Album{}
		for rows.Next() {
			var a album.Album
			if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price, &a.Available); err != nil {
				log.Fatal(err)
			}
			records = append(records, a)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(records)
	}
}

func CreateAlbum(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record album.Album
		json.NewDecoder(r.Body).Decode(&record)
		log.Println("Creating New Album...")
		err := db.QueryRow("INSERT INTO records (title, artist, price, available) VALUES ($1, $2, $3, $4) RETURNING id", record.Title, record.Artist, record.Price, record.Available).Scan(&record.Artist)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Album Created!")

		json.NewEncoder(w).Encode(record)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking Health...")
	json.NewEncoder(w).Encode("Pong")
}
