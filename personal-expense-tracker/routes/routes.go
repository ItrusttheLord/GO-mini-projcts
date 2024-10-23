package routes

import (
	"github.com/gorilla/mux"
	"personal-expense-tracker/controllers"
)

var RegisterExpensesRoutes = func(router *mux.Router) {
	router.HandleFunc("/expenses", controllers.CreateExpense).Methods("POST")
	router.HandleFunc("/expenses", controllers.GetExpenses).Methods("GET")
	router.HandleFunc("/expenses/{id}", controllers.GetExpensesById).Methods("GET")
	router.HandleFunc("/expenses/{id}", controllers.UpdateExpense).Methods("PUT")
	router.HandleFunc("/expenses/{id}", controllers.DeleteExpense).Methods("DELETE")
}
