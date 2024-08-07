package routers

import (
	"gis/internal/album"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/albums", album.GetAlbums).Methods("GET")
	router.HandleFunc("/albums", album.CreateAlbum).Methods("POST")
	router.HandleFunc("/albums/{id}", album.UpdateAlbum).Methods("PUT")
	return router
}
