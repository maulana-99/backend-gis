package routers

import (
	"gis/internal/album"
	"gis/internal/mcdonal"
	"gis/internal/users"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Albums routes
	router.HandleFunc("/albums", album.GetAlbums).Methods("GET")
	router.HandleFunc("/albums", album.CreateAlbum).Methods("POST")
	router.HandleFunc("/albums/{id}", album.UpdateAlbum).Methods("PUT")

	// McDonald's routes
	router.HandleFunc("/mcdonalds", mcdonal.GetMcDonalds).Methods("GET")
	router.HandleFunc("/mcdonalds", mcdonal.CreateMcDonald).Methods("POST")
	router.HandleFunc("/mcdonalds/{id}", mcdonal.GetMcDonaldById).Methods("GET")
	router.HandleFunc("/mcdonalds/{id}", mcdonal.UpdateMcDonald).Methods("PUT")
	router.HandleFunc("/mcdonalds/{id}", mcdonal.DeleteMcDonald).Methods("DELETE")

	// User login route
	router.HandleFunc("/users/login", users.Login).Methods("POST")
	router.HandleFunc("/users/logout", users.Logout).Methods("POST")

	// CORS preflight route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			return
		}
	}).Methods("OPTIONS")

	return router
}
