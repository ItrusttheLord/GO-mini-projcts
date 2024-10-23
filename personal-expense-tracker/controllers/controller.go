package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"personal-expense-tracker/models"
	"strconv"
	"time"
)

var Expense = []models.Expenses{
	{ID: "1", Date: time.Now(), Information: &models.Information{Name: "Coca Cola", Category: "Drink", Description: "Popular Soda drink"}, Price: &models.Price{Amount: 3.99}},
	{ID: "2", Date: time.Now(), Information: &models.Information{Name: "Apple", Category: "Fruit", Description: "Red/Green apple delicious"}, Price: &models.Price{Amount: 1.99}},
	{ID: "3", Date: time.Now(), Information: &models.Information{Name: "Sweater", Category: "Clothe", Description: "Red sweater"}, Price: &models.Price{Amount: 50.00}},
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Expense)
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	var exp models.Expenses
	w.Header().Set("Content-Type", "application/json")
	// Decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(&exp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}
	// Validate the expense
	if err := validateExpense(exp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	// Generate ID and append the expense
	ID := strconv.Itoa(rand.Intn(999999999))
	exp.ID = ID
	Expense = append(Expense, exp)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exp)
}

func GetExpensesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, ID := range Expense {
		if ID.ID == params["id"] {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(ID)
			return
		}
	}
	w.WriteHeader(404)
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, exp := range Expense {
		if exp.ID == params["id"] {
			Expense = append(Expense[:i], Expense[i+1:]...)
			w.WriteHeader(204)
			return
		}
	}
	w.WriteHeader(404)
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updateExp models.Expenses
	// Decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(&updateExp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}
	// Validate the updated expense
	if err := validateExpense(updateExp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	for i, exp := range Expense {
		if exp.ID == params["id"] {
			updateExp.ID = exp.ID
			Expense[i] = updateExp
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updateExp)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// helper func
func validateExpense(exp models.Expenses) error {
	if exp.Information.Name == "" {
		return errors.New("name cannot be empty")
	}
	if exp.Information.Category == "" {
		return errors.New("category cannot be empty")
	}
	if exp.Information.Description == "" {
		return errors.New("description cannot be empty")
	}
	if exp.Price.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}
