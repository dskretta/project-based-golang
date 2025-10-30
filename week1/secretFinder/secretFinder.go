package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// File Path
	fmt.Println("Enter the file path!")
	var filePath string
	fmt.Scanln(&filePath)

	// Keyword to search for
	fmt.Println("Enter the keyword we are searching for")
	var keyword string
	fmt.Scanln(&keyword)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Ensure it closes the file
	defer file.Close()

	// Create counters
	var lineNumber int
	lineNumber = 0
	var occurrences int
	occurrences = 0

	// Create a scanner to read from file line by line
	scanner := bufio.NewScanner(file)

	// Loop over each line in the file
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		// Check for our keyword
		if strings.Contains(line, keyword) {
			occurrences++
			fmt.Println("The keyword was found on line", lineNumber, ":", line)
		}
	}
	// After the loop is finished, print the final occurrence count
	fmt.Println("----------------------------------")
	fmt.Println("The keyword was found", occurrences, "times in file", filePath)
}
