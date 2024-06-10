package calculator

import (
	"fmt"
	"math"
	"time"
)

// CalculateArithmeticOperation returns arithmetic operation
// It starts with a capital letter, making it a public function accessible from other packages
func CalculateArithmeticOperation(operation string, a, b float64) float64 {
	// Switch operation string and calculate result
	switch operation {
	case "^":
		return math.Pow(a, b)
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		// Throw error for unknown operation
		panic("Unknown arithmetic operation, expected one of: ^ + - * /")
	}
}

// CalculateArithmeticExpression returns calculated arithmetic expression
// It starts with a capital letter, making it a public function accessible from other packages
func CalculateArithmeticExpression(expressionString string) float64 {
	startTime := time.Now()
	// Get expression as an array of numbers and operations
	ex := getExpression(expressionString)

	// Convert an expression to postfix
	ex = infixToPostfix(ex)

	// Initiate result stack array
	stack := make([]float64, 0)
	for _, value := range ex {
		stackLen := len(stack)
		switch value {
		case "+":
			if stackLen < 2 {
				panic("invalid expression: not enough operands for +")
			}
			operand2 := stack[stackLen-1]
			operand1 := stack[stackLen-2]
			stack = stack[:stackLen-2]
			stack = append(stack, operand1+operand2)
		case "-":
			if stackLen < 2 {
				panic("invalid expression: not enough operands for -")
			}
			operand2 := stack[stackLen-1]
			operand1 := stack[stackLen-2]
			stack = stack[:stackLen-2]
			stack = append(stack, operand1-operand2)
		case "*":
			if stackLen < 2 {
				panic("invalid expression: not enough operands for *")
			}
			operand2 := stack[stackLen-1]
			operand1 := stack[stackLen-2]
			stack = stack[:stackLen-2]
			stack = append(stack, operand1*operand2)
		case "/":
			if stackLen < 2 {
				panic("invalid expression: not enough operands for /")
			}
			operand2 := stack[stackLen-1]
			operand1 := stack[stackLen-2]
			stack = stack[:stackLen-2]
			stack = append(stack, operand1/operand2)
		case "^":
			if stackLen < 2 {
				panic("invalid expression: not enough operands for ^")
			}
			operand2 := stack[stackLen-1]
			operand1 := stack[stackLen-2]
			stack = stack[:stackLen-2]
			stack = append(stack, math.Pow(operand1, operand2))
		default:
			stack = append(stack, value.(float64))
		}
	}

	// Check if stack length is equal 1
	if len(stack) != 1 {
		panic("invalid expression: leftover operands")
	}

	took := time.Now().UnixNano() - startTime.UnixNano()
	fmt.Println("Finished calculating, took: ", took)
	// Return result
	return stack[0]
}

// infixToPostfix returns postfix expression
func infixToPostfix(ex expression) (result expression) {
	// Declare postfix result expression
	var stack expression

	// Iterate each value from expression
	for _, value := range ex {
		switch value {
		case "(":
			// If the value is '(', append it to the stack
			stack = append(stack, value)

		case ")":
			// If the value is ')', append stack from the end until first `(` to result
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				lastIdx := len(stack) - 1
				result = append(result, stack[lastIdx])
				stack = stack[:lastIdx]
			}
			// Check if stack length is greater then 0
			if len(stack) == 0 {
				panic("Invalid expression, mismatched brackets")
			}
			// Remove '(' character from the stack
			stack = stack[:len(stack)-1]

		case "^", "*", "/", "+", "-":
			// If the value is an operator, append stack from the end to result
			for len(stack) > 0 && getPrecedence(stack[len(stack)-1]) >= getPrecedence(value) {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, value)

		default:
			// If value is a number, append it to the result
			result = append(result, value)
		}
	}

	// Pop all the operators from the stack
	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		element := stack[lastIdx]
		if element == "(" {
			panic("Invalid expression, mismatched brackets")
		}
		result = append(result, element)
		stack = stack[:lastIdx]
	}

	// Return postfix expression
	return result
}

// getPrecedence returns precedence of operators
func getPrecedence(operation interface{}) int {
	switch operation {
	case "^":
		return 3
	case "/", "*":
		return 2
	case "+", "-":
		return 1
	default:
		return -1
	}
}
