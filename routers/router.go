package routers

import (
	"gis/internal/album"
	"gis/internal/mcdonal"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/albums", album.GetAlbums).Methods("GET")
	router.HandleFunc("/albums", album.CreateAlbum).Methods("POST")
	router.HandleFunc("/albums/{id}", album.UpdateAlbum).Methods("PUT")

	router.HandleFunc("/mcdonalds", mcdonal.GetMcDonalds).Methods("GET")
	router.HandleFunc("/mcdonalds", mcdonal.CreateMcDonald).Methods("POST")
	router.HandleFunc("/mcdonalds/{id}", mcdonal.UpdateMcDonald).Methods("PUT")
	return router
}
