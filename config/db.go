package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB adalah variabel global yang menyimpan koneksi database
var DB *sql.DB

// InitDB menginisialisasi koneksi ke database MySQL
func InitDB() {
	var err error
	// String koneksi ke database MySQL dengan username root dan database gis_db
	con := "root:@tcp(127.0.0.1:3306)/gis_db"
	// Membuka koneksi ke database dengan string koneksi yang telah ditentukan
	DB, err = sql.Open("mysql", con)
	if err != nil {
		// Menampilkan pesan error jika gagal membuka koneksi ke database
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Memeriksa apakah koneksi ke database dapat dijangkau
	if err = DB.Ping(); err != nil {
		// Menampilkan pesan error jika gagal melakukan ping ke database
		log.Fatalf("Failed to ping the database: %v", err)
	}
}

// CloseDB menutup koneksi ke database
func CloseDB() {
	// Menutup koneksi ke database dan menampilkan pesan error jika gagal
	if err := DB.Close(); err != nil {
		log.Fatalf("Failed to close the database connection: %v", err)
	}
}
