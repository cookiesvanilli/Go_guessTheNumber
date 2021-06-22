package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	counter, guess, number:= 3, -1, rand.Intn(10)
	for counter>0 && guess !=number {
		print("Guess the number 0 ...9:")
		fmt.Scanln(&guess)
		if guess !=number {
			counter --
			if guess < number {
				println("Your number is less")
			} else {
				println("Your number is greater")
			}

		}
	}
	if guess == number {
		println("You win!")
	} else {
		println("You lose: ", number)
	}
}