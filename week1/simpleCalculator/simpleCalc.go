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

	switch operator {
	case "+":
		fmt.Println(firstNum+secondNum, "is the answer")

	case "-":
		fmt.Println(firstNum-secondNum, "is the answer")

	case "*":
		fmt.Println(firstNum*secondNum, "is the answer")

	case "/":
		if secondNum == 0 {
			fmt.Println("Error: Cannot divide by zero.")
		} else {
			fmt.Println(firstNum/secondNum, "is the answer")
		}
	default:
		fmt.Println("Error: Invalid operator.")
	}
}
