package main

import (
	"gis/config"
	"gis/routers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Initialize the database
	config.InitDB()
	defer config.CloseDB()

	// Setup the router
	router := routers.SetupRouter()

	// Allow CORS
	corsHandler := cors.Default().Handler(router)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
