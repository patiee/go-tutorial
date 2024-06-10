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
		// Read `a` from standard input
		a := input.ReadFloat64("a")

		// Read `b` from standard input
		b := input.ReadFloat64("b")

		// Read arithmetic operation from standard input
		operation := input.ReadLine("operation")

		// Calculate arithmetic operation
		result := calculator.CalculateArithmeticOperation(operation, a, b)

		// Print the result as a formatted string
		// %d is a placeholder for a decimal integer
		// \n is a new line character
		fmt.Printf("%f %s %f = %f\n", a, operation, b, result)
	}
}
