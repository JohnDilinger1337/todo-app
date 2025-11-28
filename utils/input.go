package utils

import (
	"bufio"
	"fmt"
	"main/models"
	"os"
	"strconv"
	"strings"
)

// PromptText reads a string input from the user
func PromptText(msg string) string {
	fmt.Print(msg)

	// create a scanner for stdin
	scanner := bufio.NewScanner(os.Stdin)

	// read input
	scanner.Scan()

	input := scanner.Text()

	// optionally trim spaces
	input = strings.TrimSpace(input)

	// return the input
	return input
}

// PromptInt reads an integer input from the user
func PromptInt(msg string) int {
	for {
		fmt.Print(msg)

		// read line (same as PromptText)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// try to convert input to integer
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a number")
			continue // ask again
		}

		// if conversion succeeded, return value
		return value
	}
}

func PromptCategory(msg string) models.Category {
	for {
		input := PromptText(msg)

		// convert input to Category type
		category := models.Category(input)

		// check if category is valid
		if category.IsValid() {
			return category
		}
		fmt.Println("Invalid category. Try again.")
	}
}
