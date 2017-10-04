//Author Cathal Ryan
package main

import "fmt"
import "math/rand"
import "time"

//this generates random number between the range given
//Adopted this function from https://flaviocopes.com/go-random/
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main () {
    
    myrand := xrand(1,10)
	userGuess := 0
	var guess int

	fmt.Println("Guess a number between 1 and 10")

	for userGuess <10 { //For loop which asks the user to guess a number and takes in that value
		fmt.Println("Guess a number!")
		fmt.Scanf("%d", &guess)
		userGuess++

		if guess < myrand{ //if statement for if the user's guess is too low
			fmt.Println("Your guess is too low")
		}

		if guess > myrand { // if statement for if the users guess is too high
			fmt.Println("Your guess is too high")
		}
		
		if guess== myrand { // if statement for if the user has chosen the correct the number and prints the amount of tries it took.
			fmt.Printf("Correct!, That is the number, It took you %d tries\n",userGuess)
			break
		}

		

		

		
	}
}