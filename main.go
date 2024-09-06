package main

import (
	"gis/config"
	"gis/routers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Inisialisasi database saat aplikasi mulai
	config.InitDB()
	defer config.CloseDB() // Menutup koneksi database saat aplikasi berhenti

	// Mengatur router dengan rute yang telah dikonfigurasi
	router := routers.SetupRouter()

	// Mengatur CORS (Cross-Origin Resource Sharing)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Mengizinkan semua asal
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Mengizinkan metode HTTP tertentu
		AllowedHeaders: []string{"Content-Type", "Accept", "Origin", "X-Requested-With"}, // Mengizinkan header tertentu
	}).Handler(router) // Membungkus router dengan handler CORS

	// Memulai server HTTP pada port 8080
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler)) // Menghentikan aplikasi jika server mengalami error
}
