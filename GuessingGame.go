package main

import "fmt"
import "math/rand"
import "time"

//this generates random number between the range given
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main () {
    var name string
    myrand := xrand(1,6)
	userGuess := 0
	var guess int

	fmt.Println("Guess a number between 1 and 6.")

	for userGuess <6 {
		fmt.Println("Guess a number!")
		fmt.Scanf("%d", &guess)
		userGuess++

		if guess < myrand{
			fmt.Println("Your guess is too low")
		}
	}
}