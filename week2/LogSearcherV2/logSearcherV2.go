package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	// Declare variables and get User Input
	fmt.Println("Enter the directory path!")
	var dirPath string
	fmt.Scanln(&dirPath)

	fmt.Println("Enter the keyword we are searching for!")
	var keyword string
	fmt.Scanln(&keyword)

	// Read the directory contents
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for_, entry := range entries {
		fileName := entry.Name() // get's the file name

		if !entry.IsDir() {
			//It's a file!
			//Search logic
		} else {
			//It's a directory so skip?
			fmt.Println(dirPath+"/"entry.Name, " is a directory, skipping for now")
		}
	}

}
