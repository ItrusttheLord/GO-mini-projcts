package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Currency struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

const API_KEY = "ef6474b07ccd4709b0c7bf73fc2c6569"

func GetExchangeRate(base, target string) (float64, error) {
	// Build the API URL
	url := fmt.Sprintf("https://api.currencyfreaks.com/latest?apikey=%v&base=%v", API_KEY, base)

	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch exchange rate: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP status codes
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %v", err)
	}

	// Print the raw response for debugging
	fmt.Println("Raw Response Body:", string(body))

	// Decode the JSON response
	var curr Currency
	if err := json.Unmarshal(body, &curr); err != nil {
		return 0, fmt.Errorf("failed to decode response: %v", err)
	}

	// Look for the target currency in the rates
	rate, ok := curr.Rates[target]
	if !ok {
		return 0, fmt.Errorf("invalid target currency: %v", target)
	}

	// Return the exchange rate and no error
	return rate, nil
}
