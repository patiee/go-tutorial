package main

import (
	"fmt"
)

func main() {
	// A new line character is automatically passed at the end of text to `fmt.Println`
	fmt.Println("Calculating...")
	// This code prints the text '1 + 1 =' followed by the result of the arithmetic operation 1 + 1
	fmt.Println("1 + 1 =", 1+1)

	// Declare variables `a` and `b`
	// Option 1: Short variable declaration
	a, b := 2, 3
	// Option 2: Explicit variable declaration with initialization
	// var a, b int = 2, 3
	// Option 3: Explicit variable declaration without initialization, followed by assignment
	// var a, b int
	// a = 2
	// b = 3

	// Print the result as a formatted string
	// %d is a placeholder for a decimal integer
	// \n is a new line character
	// Full documentation about formatting variables you can find here: https://pkg.go.dev/fmt
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}
