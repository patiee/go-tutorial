package calculator

import (
	"fmt"
	"strconv"
)

// expression is a custom type representing an array of interface{}.
// In this array, float64 values represent numbers, while string values represent operations.
type expression []interface{}

// getExpression returns an array of numbers and operations
func getExpression(expressionString string) (ex expression) {
	// Declare variable for string number
	var value string

	// Create an loop to read every character from `expressionString`
	// Each character will be read as rune to `ch`
	// `_` is used to omit and not declare value for unused value, here is index number from `expressionString`
	for _, ch := range expressionString {
		switch ch {
		case ' ':
			continue
		case '(', ')', '^', '+', '-', '*', '/':
			// Check if number value is non-empty
			if value != "" {
				// Convert string `value` number into float
				number, err := strconv.ParseFloat(value, 64)
				if err != nil {
					panic(fmt.Sprintf("Could not parse %q into float64", value))
				}

				// Append to expression a number
				ex = append(ex, number)
				// Clear value
				value = ""
			}
			// Append to expression an operation
			ex = append(ex, string(ch))

		default:
			// read number
			// add string from `ch` to value
			value += string(ch)
		}
	}

	// Check if there's non empty value
	if value != "" {
		// Convert string `value` number into float
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(fmt.Sprintf("Could not parse %q into float64", value))
		}

		// Append to expression a number
		ex = append(ex, number)
	}

	// Return an expression
	return ex
}
