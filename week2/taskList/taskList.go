package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// This slice holds all the tasks
	var tasks []string

	// This is the main program loop
	for {
		//  Print the menu
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1. Add a new task")
		fmt.Println("2. List all tasks")
		fmt.Println("3. Quit")

		// Get user's choice
		var userInput int
		var taskNumber = 0
		fmt.Scanln(&userInput)

		// Switch to choose
		switch userInput {
		case 1:
			// Add logic
			fmt.Println("Enter the new task:")

			// Use bufio.NewReader to read the whole line
			reader := bufio.NewReader(os.Stdin)
			newTask, err := reader.ReadString('\n')

			// Error checking
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			// Actual handling of tasks
			newTask = strings.TrimSpace(newTask)

			// Add the clean task to our slice
			if newTask != "" { // doesn't add empty tasks
				tasks = append(tasks, newTask)
				taskNumber++
				fmt.Println("Task added!")
			}

		case 2:
			// List tasks logic
			fmt.Print("\n You have ", taskNumber, " tasks", "\n--- Here are all your tasks: ---")

			// For range loop to print neatly
			for _, value := range tasks {
				fmt.Println("\n- " + value) // Add a dash to each
			}

			fmt.Println("--------------------------------")

		case 3:
			// Quit logic
			fmt.Println("Goodbye!")
			return // Return 'exits' func main

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}

}
