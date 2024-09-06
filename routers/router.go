package routers

import (
	"gis/internal/album"
	"gis/internal/mcdonal"
	"gis/internal/users"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouter mengatur semua route untuk aplikasi dan mengembalikan router yang sudah dikonfigurasi
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Routes untuk Albums
	router.HandleFunc("/albums", album.GetAlbums).Methods("GET")      // Mengambil semua album
	router.HandleFunc("/albums", album.CreateAlbum).Methods("POST") // Menambahkan album baru
	router.HandleFunc("/albums/{id}", album.UpdateAlbum).Methods("PUT") // Memperbarui album berdasarkan ID

	// Routes untuk McDonald's
	router.HandleFunc("/mcdonalds", mcdonal.GetMcDonalds).Methods("GET")     // Mengambil semua data McDonald's
	router.HandleFunc("/mcdonalds", mcdonal.CreateMcDonald).Methods("POST")  // Menambahkan McDonald's baru
	router.HandleFunc("/mcdonalds/{id}", mcdonal.GetMcDonaldById).Methods("GET") // Mengambil McDonald's berdasarkan ID
	router.HandleFunc("/mcdonalds/{id}", mcdonal.UpdateMcDonald).Methods("PUT")  // Memperbarui McDonald's berdasarkan ID
	router.HandleFunc("/mcdonalds/{id}", mcdonal.DeleteMcDonald).Methods("DELETE") // Menghapus McDonald's berdasarkan ID

	// Routes untuk login dan logout pengguna
	router.HandleFunc("/users/login", users.Login).Methods("POST") // Login pengguna
	router.HandleFunc("/users/logout", users.Logout).Methods("POST") // Logout pengguna

	// Route untuk CORS preflight
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Menetapkan header CORS untuk memungkinkan permintaan dari semua asal
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			// Mengembalikan response untuk preflight OPTIONS request
			return
		}
	}).Methods("OPTIONS")

	return router
}
