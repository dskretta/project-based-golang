package main

import (
	"fmt"
)

func main() {
	// 1. Get the first number
	fmt.Println("What is the first variable in the equation?")
	var firstNum float64
	fmt.Scanln(&firstNum)

	// 2. Get the operator
	fmt.Println("Enter the operator (+, -, *, /):")
	var operator string
	fmt.Scanln(&operator)

	// 3. Get the second number
	fmt.Println("What is the second variable in the equation?")
	var secondNum float64
	fmt.Scanln(&secondNum)

	// 4. Find out the operator
	if operator == "+" {
		fmt.Println(firstNum+secondNum, "is the answer")
	} else if operator == "-" {
		fmt.Println(firstNum-secondNum, "is the answer")
	} else if operator == "*" {
		fmt.Println(firstNum*secondNum, "is the answer")
	} else if operator == "/" {
		// 5. Check if secondNum is 0 to avoid a crash
		if secondNum == 0 {
			fmt.Println("Error: Cannot divide by zero.")
		} else {
			// If it's not 0, it's safe to divide
			fmt.Println(firstNum/secondNum, "is the answer")
		}
	} else {
		// 6. Add a final "else" to catch bad input
		fmt.Println("Error: Invalid operator.")
	}
}
