package mcdonal

import (
	"encoding/json"
	"gis/config"
	"gis/internal/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMcDonalds(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, latitude, longitude FROM mcdonals")
	if err != nil {
		log.Printf("Error querying the database: %v", err)
		http.Error(w, "Error querying the database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var mcdonalds []models.Mcdonald
	for rows.Next() {
		var mcdonald models.Mcdonald
		if err := rows.Scan(&mcdonald.ID, &mcdonald.Name, &mcdonald.Latitude, &mcdonald.Longitude); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		mcdonalds = append(mcdonalds, mcdonald)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		http.Error(w, "Error iterating rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetMcDonaldById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	var mcdonald models.Mcdonald
	err := config.DB.QueryRow("SELECT id, name, latitude, longitude FROM mcdonals WHERE id = ?", id).Scan(&mcdonald.ID, &mcdonald.Name, &mcdonald.Latitude, &mcdonald.Longitude)
	if err != nil {
		log.Printf("Error getting McDonald by id: %v", err)
		http.Error(w, "Error getting McDonald by id: "+err.Error(), http.StatusInternalServerError)
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateMcDonald(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var mcdonald models.Mcdonald
	if err := json.NewDecoder(r.Body).Decode(&mcdonald); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Use positional parameters instead of named parameters
	_, err := config.DB.Exec("INSERT INTO mcdonals (name, latitude, longitude) VALUES (?, ?, ?)",
		mcdonald.Name, mcdonald.Latitude, mcdonald.Longitude)
	if err != nil {
		log.Printf("Error inserting McDonald: %v", err)
		http.Error(w, "Error inserting McDonald: "+err.Error(), http.StatusInternalServerError)
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

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
	err := config.DB.QueryRow("SELECT id, name, latitude, longitude FROM mcdonals WHERE id = ?", id).Scan(&mcdonald.ID, &mcdonald.Name, &mcdonald.Latitude, &mcdonald.Longitude)
	if err != nil {
		log.Printf("Error getting McDonald by id for update: %v", err)
		http.Error(w, "Error getting McDonald by id for update: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&mcdonald); err != nil {
		log.Printf("Error decoding JSON for update: %v", err)
		http.Error(w, "Error decoding JSON for update: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("UPDATE mcdonals SET name = ?, latitude = ?, longitude = ? WHERE id = ?", mcdonald.Name, mcdonald.Latitude, mcdonald.Longitude, id)
	if err != nil {
		log.Printf("Error updating McDonald: %v", err)
		http.Error(w, "Error updating McDonald: "+err.Error(), http.StatusInternalServerError)
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

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

	_, err := config.DB.Exec("DELETE FROM mcdonals WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting McDonald: %v", err)
		http.Error(w, "Error deleting McDonald: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := EmptyResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "McDonald deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
