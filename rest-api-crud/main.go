package main

import (
	"log"
	"net/http"
	"rest-api-crud/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterAnimeRoutes(router)

	log.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
