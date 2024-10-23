package main

import (
	"URLshortener/middleware"
	"URLshortener/routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	logger := middlware.NewLogger(router)
	routes.GetUrlRoutes(router)
	fmt.Println("Starting Route on port :8080")
	log.Fatal(http.ListenAndServe(":8080", logger))
}
