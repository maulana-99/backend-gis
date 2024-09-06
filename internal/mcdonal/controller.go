package mcdonal

import (
	"encoding/json"
	"gis/config"
	"gis/internal/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetMcDonalds menangani permintaan HTTP GET untuk mengambil semua data McDonald's dari database
func GetMcDonalds(w http.ResponseWriter, r *http.Request) {
	// Menjalankan query untuk mengambil data McDonald's dari database
	rows, err := config.DB.Query("SELECT id, name, latitude, longitude FROM mcdonals")
	if err != nil {
		log.Printf("Error querying the database: %v", err)
		http.Error(w, "Error querying the database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var mcdonalds []models.Mcdonald
	// Iterasi hasil query dan memindahkan data ke dalam slice mcdonalds
	for rows.Next() {
		var mcdonald models.Mcdonald
		if err := rows.Scan(&mcdonald.ID, &mcdonald.Name, &mcdonald.Latitude, &mcdonald.Longitude); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		mcdonalds = append(mcdonalds, mcdonald)
	}

	// Memeriksa apakah ada kesalahan saat iterasi hasil query
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		http.Error(w, "Error iterating rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Menyiapkan response dengan data McDonald's
	response := EmptyResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "McDonalds retrieved successfully",
		Data: struct {
			McDonalds []models.Mcdonald `json:"mcdonalds"`
		}{
			McDonalds: mcdonalds,
		},
	}

	// Mengatur header response dan mengirim data dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetMcDonaldById menangani permintaan HTTP GET untuk mengambil data McDonald's berdasarkan ID
func GetMcDonaldById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	var mcdonald models.Mcdonald
	// Menjalankan query untuk mengambil data McDonald's berdasarkan ID
	err := config.DB.QueryRow("SELECT id, name, latitude, longitude FROM mcdonals WHERE id = ?", id).Scan(&mcdonald.ID, &mcdonald.Name, &mcdonald.Latitude, &mcdonald.Longitude)
	if err != nil {
		log.Printf("Error getting McDonald by id: %v", err)
		http.Error(w, "Error getting McDonald by id: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Menyiapkan response dengan data McDonald's
	response := EmptyResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "McDonald retrieved successfully",
		Data: struct {
			McDonalds []models.Mcdonald `json:"mcdonalds"`
		}{
			McDonalds: []models.Mcdonald{mcdonald},
		},
	}

	// Mengatur header response dan mengirim data dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateMcDonald menangani permintaan HTTP POST untuk menambahkan data McDonald's baru
func CreateMcDonald(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var mcdonald models.Mcdonald
	// Mendecode body request menjadi objek McDonald's
	if err := json.NewDecoder(r.Body).Decode(&mcdonald); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Menjalankan query untuk menambahkan data McDonald's ke database
	_, err := config.DB.Exec("INSERT INTO mcdonals (name, latitude, longitude) VALUES (?, ?, ?)",
		mcdonald.Name, mcdonald.Latitude, mcdonald.Longitude)
	if err != nil {
		log.Printf("Error inserting McDonald: %v", err)
		http.Error(w, "Error inserting McDonald: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Menyiapkan response untuk konfirmasi bahwa data McDonald's telah ditambahkan
	response := EmptyResponse{
		Success: true,
		Code:    http.StatusCreated,
		Message: "McDonald created successfully",
		Data: struct {
			McDonalds []models.Mcdonald `json:"mcdonalds"`
		}{
			McDonalds: []models.Mcdonald{mcdonald},
		},
	}

	// Mengatur header response dan mengirim data dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateMcDonald menangani permintaan HTTP PUT untuk memperbarui data McDonald's berdasarkan ID
func UpdateMcDonald(w http.ResponseWriter, r *http.Request) {
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

	var mcdonald models.Mcdonald
	// Menjalankan query untuk mengambil data McDonald's berdasarkan ID sebelum diperbarui
	err := config.DB.QueryRow("SELECT id, name, latitude, longitude FROM mcdonals WHERE id = ?", id).Scan(&mcdonald.ID, &mcdonald.Name, &mcdonald.Latitude, &mcdonald.Longitude)
	if err != nil {
		log.Printf("Error getting McDonald by id for update: %v", err)
		http.Error(w, "Error getting McDonald by id for update: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Mendecode body request menjadi objek McDonald's yang baru
	if err := json.NewDecoder(r.Body).Decode(&mcdonald); err != nil {
		log.Printf("Error decoding JSON for update: %v", err)
		http.Error(w, "Error decoding JSON for update: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Menjalankan query untuk memperbarui data McDonald's di database
	_, err = config.DB.Exec("UPDATE mcdonals SET name = ?, latitude = ?, longitude = ? WHERE id = ?", mcdonald.Name, mcdonald.Latitude, mcdonald.Longitude, id)
	if err != nil {
		log.Printf("Error updating McDonald: %v", err)
		http.Error(w, "Error updating McDonald: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Menyiapkan response untuk konfirmasi bahwa data McDonald's telah diperbarui
	response := EmptyResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "McDonald updated successfully",
		Data: struct {
			McDonalds []models.Mcdonald `json:"mcdonalds"`
		}{
			McDonalds: []models.Mcdonald{mcdonald},
		},
	}

	// Mengatur header response dan mengirim data dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteMcDonald menangani permintaan HTTP DELETE untuk menghapus data McDonald's berdasarkan ID
func DeleteMcDonald(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// Menjalankan query untuk menghapus data McDonald's berdasarkan ID
	_, err := config.DB.Exec("DELETE FROM mcdonals WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting McDonald: %v", err)
		http.Error(w, "Error deleting McDonald: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Menyiapkan response untuk konfirmasi bahwa data McDonald's telah dihapus
	response := EmptyResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "McDonald deleted successfully",
	}

	// Mengatur header response dan mengirim data dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
