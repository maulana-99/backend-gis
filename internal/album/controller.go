package album

import (
	"encoding/json"
	"gis/config"
	"gis/internal/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	response := EmptyResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "Albums retrieved successfully",
		Data: struct {
			Albums []models.Album `json:"albums"`
		}{
			Albums: albums,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

	response := EmptyResponse{
		Success: true,
		Code:    http.StatusCreated,
		Message: "Album created successfully",
		Data: struct {
			Albums []models.Album `json:"albums"`
		}{
			Albums: []models.Album{album},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	var album models.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := config.DB.Exec("UPDATE albums SET title = ?, artist = ?, price = ? WHERE id = ?", album.Title, album.Artist, album.Price, id)
	if err != nil {
		log.Printf("Error updating album: %v", err)
		http.Error(w, "Error updating album: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := EmptyResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "Album updated successfully",
		Data: struct {
			Albums []models.Album `json:"albums"`
		}{
			Albums: []models.Album{album},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
