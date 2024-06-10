package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readInt reads a line of text from the provided reader and converts it to an int
func readInt(reader *bufio.Reader, name string) int {
	// Print a prompt asking the user to enter a value for the given name
	fmt.Printf("Enter %s: ", name)

	// Read text from reader until new line character `\n`
	valueString, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("error while reading a: %s", err))
	}

	// Remove additional characters from text like space, new line
	valueString = strings.TrimSpace(valueString)
	// Parse input text to an int variable
	value, err := strconv.Atoi(valueString)
	if err != nil {
		panic(fmt.Sprintf("could not parse a into int: %s", err))
	}

	// Return value from reader as int
	return value
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

		// Print the result as a formatted string
		// %d is a placeholder for a decimal integer
		// \n is a new line character
		fmt.Printf("%d + %d = %d\n", a, b, a+b)
	}
}
