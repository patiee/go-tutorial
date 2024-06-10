package main

import (
	"fmt"

	"github.com/patiee/go-tutorial/calculator"
	"github.com/patiee/go-tutorial/input"
)

func main() {
	// Create an infinite loop for input
	// You can exit program with `control + C`
	for {
		// Read arithmetic expression from standard input, example: `6 + 2 * ( 3 - 1 ) ^ 2`
		expression := input.ReadLine("arithmetic expression")

		// Calculate arithmetic operation
		result := calculator.CalculateArithmeticExpression(expression)

		// Print the result as a formatted string
		// %d is a placeholder for a decimal integer
		// \n is a new line character
		fmt.Printf("= %f\n", result)
	}
}
