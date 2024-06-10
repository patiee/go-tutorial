package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// readInt reads a line of text from the provided reader and converts it to an int
func readInt(reader *bufio.Reader, name string) int {
	// Read line from reader
	valueString := readLine(reader, name)

	// Parse input text to an int variable
	value, err := strconv.Atoi(valueString)
	if err != nil {
		panic(fmt.Sprintf("could not parse a into int: %s", err))
	}

	// Return value from reader as int
	return value
}

// readLine reads a line of text from the provided reader and returns string
func readLine(reader *bufio.Reader, name string) string {
	// Print a prompt asking the user to enter a value for the given name
	fmt.Printf("Enter %s: ", name)

	// Read text from reader until new line character `\n`
	valueString, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("error while reading a: %s", err))
	}

	// Remove additional characters from text like space, new line
	valueString = strings.TrimSpace(valueString)
	// Return value from reader
	return valueString
}

func main() {
	// Define a reader to read text bytes from os.Stdin which is an input from terminal to program
	reader := bufio.NewReader(os.Stdin)

	// Create an infinite loop for input
	// You can exit program with `control + C`
	for {
		// Read `a` from reader
		a := readInt(reader, "a")

		// Read `b` from reader
		b := readInt(reader, "b")

		// Read arithmetic operation
		operation := readLine(reader, "operation")

		// Declare result variable
		var result int

		// Switch operation string and calculate result
		switch operation {
		case "^":
			result = int(math.Pow(float64(a), float64(b)))
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			// Integer result doesn't have decimal precision for division
			// Let's change `int` to `float64` in the next example
			result = a / b
		default:
			// Throw error for unknown operation
			panic("Unknown arithmetic operation, expected one of: + - * /")
		}

		// Print the result as a formatted string
		// %d is a placeholder for a decimal integer
		// \n is a new line character
		fmt.Printf("%d %s %d = %d\n", a, operation, b, result)
	}
}
