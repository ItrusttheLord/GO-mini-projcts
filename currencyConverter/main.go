package main

import (
	"bufio"
	"currencyConverter/api"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ask for base currency
	fmt.Print("Please enter the base currency (e.g., USD): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	baseCurr := scanner.Text()
	// Ask for target currency
	fmt.Print("Please enter the target currency (e.g., EUR): ")
	scanner.Scan()
	targetCurr := scanner.Text()
	// Ask for the amount to convert
	fmt.Print("Please enter the amount you want to convert: ")
	scanner.Scan()
	amountStr := scanner.Text()
	// Convert the amount to a float
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount entered")
		return
	}
	// Fetch the exchange rate
	exchangeRate, err := api.GetExchangeRate(baseCurr, targetCurr)
	if err != nil {
		fmt.Printf("Error fetching exchange rate: %v\n", err)
		return
	}
	// Calculate the converted amount
	convertedAmount := amount * exchangeRate
	// Print the result
	fmt.Printf("%.2f %s is equivalent to %.2f %s\n", amount, baseCurr, convertedAmount, targetCurr)
}
