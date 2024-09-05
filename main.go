package main

import (
	"gis/config"
	"gis/routers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Initialize the database once when the application starts
	config.InitDB()
	defer config.CloseDB() // Close the database when the application stops

	// Setup the router
	router := routers.SetupRouter()

	// Allow CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Origin", "X-Requested-With"},
	}).Handler(router)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
