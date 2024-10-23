package routes

import (
	"github.com/gorilla/mux"
	"rest-api-crud/controllers"
)

var RegisterAnimeRoutes = func(router *mux.Router) {
	router.HandleFunc("/animes", controllers.CreateAnime).Methods("POST")
	router.HandleFunc("/animes", controllers.GetAnimes).Methods("GET")
	router.HandleFunc("/animes/{id}", controllers.GetAnimeById).Methods("GET")
	router.HandleFunc("/animes/{id}", controllers.UpdateAnime).Methods("PUT")
	router.HandleFunc("/animes/{id}", controllers.DeleteAnime).Methods("DELETE")
}
