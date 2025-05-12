package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Welcome to the guessing game")
	fmt.Println("rando number between 1-100, start guessing")

	var guess int

	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Correct!")
			break
		} else if guess > target {
			fmt.Println("Too high!")
			continue
		} else if guess < target {
			fmt.Println("Too low!")
			continue
		}
	}

}
