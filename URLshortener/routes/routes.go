package routes

import (
	"URLshortener/controllers"
	"github.com/gorilla/mux"
)

var GetUrlRoutes = func(router *mux.Router) {
	router.HandleFunc("/post", controllers.HandlePost).Methods("POST")
	router.HandleFunc("/get", controllers.HandleGet).Methods("GET")
	router.HandleFunc("/getlist", controllers.HandleList).Methods("GET")
	router.HandleFunc("/delete", controllers.HandleDelete).Methods("delete")
}
