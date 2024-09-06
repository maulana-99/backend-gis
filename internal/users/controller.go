package users

import (
	"encoding/json"
	"gis/config"
	"gis/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Key untuk menandatangani JWT
var jwtKey = []byte("uajsbbkz")

// Login menangani login pengguna
func Login(w http.ResponseWriter, r *http.Request) {
	// Struktur untuk menampung username dan password dari request body
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Decode JSON dari request body ke struct credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Query untuk mengambil pengguna berdasarkan username dan password yang diberikan
	rows, err := config.DB.Query("SELECT id, username, password FROM users WHERE username = ? AND password = ?", credentials.Username, credentials.Password)
	if err != nil {
		log.Printf("Error querying the database: %v", err)
		http.Error(w, "Error querying the database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Pastikan rows ditutup setelah selesai

	var users []models.Users
	// Iterasi hasil query
	for rows.Next() {
		var user models.Users
		// Pindahkan hasil query ke dalam struct user
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user) // Tambahkan user ke dalam slice users
	}

	// Cek jika ada error saat iterasi rows
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		http.Error(w, "Error iterating rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Jika tidak ada pengguna ditemukan, kirim respons Unauthorized
	if len(users) == 0 {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Membuat klaim untuk JWT
	expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku selama 24 jam
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   credentials.Username, // Set subject sebagai username
	}

	// Buat token JWT menggunakan klaim yang telah dibuat
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Error generating JWT token: %v", err)
		http.Error(w, "Error generating JWT token", http.StatusInternalServerError)
		return
	}

	// Set token JWT sebagai cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",          // Nama cookie
		Value:   tokenString,      // Nilai token
		Expires: expirationTime,   // Expiration time cookie
	})

	// Buat respons JSON untuk menunjukkan login berhasil
	response := EmptyResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "Users Login successfully",
		Data: struct {
			Users []models.Users `json:"users"`
		}{
			Users: users, // Data user yang login
		},
	}

	// Kirim respons sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Logout menangani logout pengguna
func Logout(w http.ResponseWriter, r *http.Request) {
	// Set cookie token dengan nilai kosong dan waktu kedaluwarsa di masa lalu untuk menghapusnya
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour), // Mengatur cookie agar segera kadaluwarsa
	})

	// Buat respons JSON untuk menunjukkan logout berhasil
	response := EmptyResponse{
		Success: true,
		Code:    http.StatusOK,
		Message: "User logged out successfully",
	}

	// Kirim respons sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
