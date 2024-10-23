package main

import (
	"bufio"
	"calculatorAPI/cmd"
	"calculatorAPI/middleware"
	"calculatorAPI/routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	isValid := false
	//ask the user until he enters a valid response
	for !isValid {
		fmt.Printf("Pick a mode: terminal or server: ")
		scanner.Scan()
		input := scanner.Text()
		//if user enters terminal we run the terminal
		if input == "terminal" {
			cmd.RunCalculator()
			isValid = true //if server we start the server
		} else if input == "server" {
			router := mux.NewRouter()
			logger := middleware.NewLogger(router)
			routes.GetCalculatorRoutes(router)
			fmt.Println("Starting route on port :8080")
			log.Fatal(http.ListenAndServe(":8080", logger))
			isValid = true
		} else {
			fmt.Println("Invalid Choice! Please pick terminal or server.")
		}
	}
}
