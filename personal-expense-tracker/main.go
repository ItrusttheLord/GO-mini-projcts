package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"personal-expense-tracker/routes"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterExpensesRoutes(route)
	println("Starting Route on port :8080")
	if err := http.ListenAndServe(":8080", route); err != nil {
		log.Fatal()
	}
}
