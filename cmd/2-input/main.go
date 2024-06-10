package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Define a reader to read text bytes from os.Stdin which is an input from terminal to program
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a: ")
	// Read text from reader until new line character `\n`
	aString, err := reader.ReadString('\n')
	if err != nil {
		// Panic prints an error message and exits the program immediately
		panic(fmt.Sprintf("error while reading a: %s", err))
	}

	// Remove additional characters from text like space, new line
	aString = strings.TrimSpace(aString)
	// Parse input text to an int variable
	a, err := strconv.Atoi(aString)
	if err != nil {
		panic(fmt.Sprintf("could not parse a into int: %s", err))
	}

	fmt.Print("Enter b: ")
	// Read text from reader until new line character `\n`
	bString, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("error while reading b: %s", err))
	}

	// Parse trimmed input text to b int variable
	b, err := strconv.Atoi(strings.TrimSpace(bString))
	if err != nil {
		panic(fmt.Sprintf("could not parse b into int: %s", err))
	}

	// Print result as formatted string text
	// %d is for printing number as decimal
	// \n is a new line character
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}
