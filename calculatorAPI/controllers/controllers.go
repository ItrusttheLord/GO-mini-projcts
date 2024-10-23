package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Results struct {
	Result int `json:"result"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Helper func to get the numbers from the queries
func getNumbersFromQuery(w http.ResponseWriter, r *http.Request) (int, int, error) {
	queries := r.URL.Query() // Get queries
	num1 := queries.Get("num1")
	num2 := queries.Get("num2")
	// Check if both parameters are present
	if !queries.Has("num1") || !queries.Has("num2") {
		response := Response{
			Status:  "error",
			Message: "Missing parameters: num1 and num2 are required.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return 0, 0, fmt.Errorf("missing parameters: num1 and num2")
	}
	// Convert strings to int
	int1, err1 := strconv.Atoi(num1)
	int2, err2 := strconv.Atoi(num2)
	if err1 != nil || err2 != nil {
		response := Response{
			Status:  "error",
			Message: "Invalid input: num1 and num2 must be numeric.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return 0, 0, fmt.Errorf("invalid input: num1 and num2 must be numeric")
	}
	return int1, int2, nil
}

// addition handler
func AddHandler(w http.ResponseWriter, r *http.Request) {
	num1, num2, err := getNumbersFromQuery(w, r)
	if err != nil {
		return
	}
	result := num1 + num2
	res := Results{Result: result}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// subtraction handler
func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	num1, num2, err := getNumbersFromQuery(w, r)
	if err != nil {
		return
	}
	result := num1 - num2
	res := Results{Result: result}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// multiplication handler
func MultiplicationHandler(w http.ResponseWriter, r *http.Request) {
	num1, num2, err := getNumbersFromQuery(w, r)
	if err != nil {
		return
	}
	result := num1 * num2
	res := Results{Result: result}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// division handler
func DivisionHandler(w http.ResponseWriter, r *http.Request) {
	num1, num2, err := getNumbersFromQuery(w, r)
	if err != nil {
		return
	}
	//check divisin by 0(if we don't this the server will get a panic)
	if num2 == 0 {
		response := Response{
			Status:  "error",
			Message: "Division by zero is not allowed",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	result := num1 / num2
	res := Results{Result: result}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Terminal Use
func Add(n1, n2 int) (int, error) {
	result := n1 + n2
	return result, nil
}

func Subtr(n1, n2 int) (int, error) {
	result := n1 - n2
	return result, nil
}

func Multiplication(n1, n2 int) (int, error) {
	result := n1 * n2
	return result, nil
}

func Division(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	result := n1 / n2
	return result, nil
}
