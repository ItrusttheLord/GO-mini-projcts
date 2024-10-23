package routes

import (
	"calculatorAPI/controllers"

	"github.com/gorilla/mux"
)

// create routes for the handlers
var GetCalculatorRoutes = func(router *mux.Router) {
	router.HandleFunc("/add", controllers.AddHandler).Methods("GET")
	router.HandleFunc("/subtr", controllers.SubtractHandler).Methods("GET")
	router.HandleFunc("/multi", controllers.MultiplicationHandler).Methods("GET")
	router.HandleFunc("/divis", controllers.DivisionHandler).Methods("GET")
}
