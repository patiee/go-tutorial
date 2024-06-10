package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// reader is a private global variable for the package `input`
// It starts with a lowercase letter, making it a variable accessible only from `input` package
// This Reader is initialized to read text bytes from os.Stdin, which represents input from the terminal to the program
var reader = bufio.NewReader(os.Stdin)

// ReadFloat64 reads a line of text from the provided reader and converts it to float64
// It starts with a capital letter, making it a public function accessible from other packages
func ReadFloat64(name string) float64 {
	// Read line from reader
	valueString := ReadLine(name)

	// Parse input text to float64 variable
	// BitSize is 64 for float64 and 32 for float32
	value, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		panic(fmt.Sprintf("could not parse a into float64: %s", err))
	}

	// Return value from reader as int
	return value
}

// ReadLine reads a line of text from the provided reader and returns string
// It starts with a capital letter, making it a public function accessible from other packages
func ReadLine(name string) string {
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
