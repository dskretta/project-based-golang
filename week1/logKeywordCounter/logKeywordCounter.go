package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Prompt the user for filePath
	fmt.Println("Enter the file path:")
	var filePath string
	fmt.Scanln(&filePath)

	// 2. Prompt the user for keyword
	fmt.Println("Enter the keyword to search for:")
	var keyword string
	fmt.Scanln(&keyword)

	// 3. Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// 4. Defer the file close
	defer file.Close()

	// 5. Declare counter variable
	var occurrences int
	occurrences = 0

	// 6. Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// 7. Loop over each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// 8. Check if the file contains our keyword
		if strings.Contains(line, keyword) {
			occurrences = occurrences + 1
		}
	}
	// 8 After the loop ends, print the final result
	fmt.Println(keyword, "occures in the file", occurrences, "times!")
}
