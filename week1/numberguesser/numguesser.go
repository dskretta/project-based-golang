package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1, Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// 2. Pick a secret number (1-100)
	secretNumber := rand.Intn(100) + 1

	// 3. Prompt the user
	fmt.Println("I've picked a random number between 1 and 100. Can you guess it?")

	// 4. Create a variable to store the guess
	var guess int

	for {
		// 5. Read the user's input
		fmt.Scanln(&guess)

		// 6. Compare the guess to the secret number
		if guess < secretNumber {
			fmt.Println("Too low! Guess again.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Guess again.")
		} else if guess == secretNumber {
			fmt.Println("You are correct! Congratulations!")
			break // Stop the loop if correct
		}

	}
}
