package cmd

import (
	"bufio"
	"calculatorAPI/controllers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Main logic
func RunCalculator() {
	scanner := bufio.NewScanner(os.Stdin) //create new scanner

	for {
		fmt.Printf("Choose an Operation (add, subtr, multi, divis)? ")
		scanner.Scan()          //read input
		input := scanner.Text() //save user input

		switch input {
		case "add":
			handleOperation(scanner, "addition", controllers.Add)
		case "subtr":
			handleOperation(scanner, "subtraction", controllers.Subtr)
		case "multi":
			handleOperation(scanner, "multiplication", controllers.Multiplication)
		case "divis":
			handleOperation(scanner, "division", controllers.Division)
		default:
			fmt.Println("Invalid operation. Please try again.")
			continue
		}

		// Ask if the user wants to continue or quit
		fmt.Printf("Would you like to continue? (y/q): ")
		scanner.Scan()
		if scanner.Text() == "q" {
			return // Exit the loop and function
		}
	}
}

// helper func to split the user input into two values
func getTwoNumbers(scanner *bufio.Scanner) (int, int, error) {
	fmt.Printf("Please Enter Two numbers (ex: 2 3): ")
	scanner.Scan()                             //scan use input
	input := strings.TrimSpace(scanner.Text()) //save the user innput
	values := strings.Split(input, " ")        //split user input
	// Validate input length
	if len(values) != 2 {
		return 0, 0, fmt.Errorf("please enter exactly two numbers")
	}
	// Parse the numbers
	num1, err1 := strconv.Atoi(values[0])
	num2, err2 := strconv.Atoi(values[1])
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("invalid input enter valid integers")
	}
	return num1, num2, nil
}

// handler operation logic
func handleOperation(scanner *bufio.Scanner, operation string, operationFunc func(int, int) (int, error)) {
	num1, num2, err := getTwoNumbers(scanner)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Call the operation function (e.g., Add, Subtr, Multiplication, Division)
	res, err := operationFunc(num1, num2)
	if err != nil {
		fmt.Printf("Error in %s: %v\n", operation, err)
		return
	}

	fmt.Printf("Result of %s: %v\n", operation, res)
}
