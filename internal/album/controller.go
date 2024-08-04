package album

import (
	"encoding/json"
	"gis/config"
	"gis/internal/models"
	"log"
	"net/http"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		log.Printf("Error querying the database: %v", err)
		http.Error(w, "Error querying the database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var albums []models.Album
	for rows.Next() {
		var album models.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		http.Error(w, "Error iterating rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var album models.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := config.DB.Exec("INSERT INTO albums (title, artist, price) VALUES (:title, :artist, :price)", album.Title, album.Artist, album.Price)
	if err != nil {
		log.Printf("Error inserting album: %v", err)
		http.Error(w, "Error inserting album: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(album)
}
