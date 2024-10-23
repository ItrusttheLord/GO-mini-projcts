package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Questions struct {
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
	Answer   string   `json:"answer"`
}

func LoadQuestions() ([]Questions, error) {
	var questions []Questions
	file, err := os.ReadFile("questions.json") //read JSON file
	if err != nil {
		return nil, fmt.Errorf("failed to read file questions.json: %w", err)
	}
	err = json.Unmarshal(file, &questions) //Unmarshan content
	if err != nil {
		return nil, fmt.Errorf("failed to decode file questions.json: %w", err)
	}
	return questions, nil
}

func main() {
	// Load the questions from JSON file
	readQuestions, err := LoadQuestions()
	userScore := 0
	if err != nil {
		log.Println(fmt.Errorf("failed to read file questions.json: %w", err))
		return
	}
	// Welcome message
	fmt.Println("Hello, welcome to the Trivia Quiz!")
	fmt.Println("Begin when you are ready!")
	fmt.Println("Please select an answer")
	// Create scanner only once
	scanner := bufio.NewScanner(os.Stdin)
	// Loop through each question
	for _, q := range readQuestions {
		fmt.Printf("%v\n", q.Question)
		// Print choices with numbering for clarity
		for index, choice := range q.Choices {
			fmt.Printf("%d. %v\n", index+1, choice) // Number choices starting from 1
		}
		var input string
		validInput := false
		// Keep prompting the user until they enter a valid answer
		for !validInput {
			fmt.Printf("Enter your answer here: ")
			scanner.Scan()
			input = scanner.Text()
			// Check if input matches any valid choice
			for _, choice := range q.Choices {
				if strings.TrimSpace(strings.ToLower(input)) == strings.TrimSpace(strings.ToLower(choice)) {
					validInput = true
					break
				}
			}
			if !validInput {
				fmt.Println("Please enter a valid answer from the given choices.")
			}
		}
		// Check if the answer is correct
		if strings.TrimSpace(strings.ToLower(input)) == strings.TrimSpace(strings.ToLower(q.Answer)) {
			fmt.Println("Correct answer!")
			userScore++
		} else {
			fmt.Printf("Incorrect answer. The correct answer is: %v\n", q.Answer)
		}
		fmt.Println() // Add a blank line for better readability between questions
	}
	// Display the final score
	fmt.Printf("Your Score is: %v/%v\n", userScore, len(readQuestions))
}
