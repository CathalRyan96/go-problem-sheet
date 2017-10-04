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
    
    myrand := xrand(1,10)
	userGuess := 0
	var guess int

	fmt.Println("Guess a number between 1 and 10")

	for userGuess <6 {
		fmt.Println("Guess a number!")
		fmt.Scanf("%d", &guess)
		userGuess++

		if guess < myrand{
			fmt.Println("Your guess is too low")
		}

		if guess > myrand {
			fmt.Println("Your guess is too high")
		}
		
		if guess== myrand {
			fmt.Printf("Correct!, That is the number, It took you %d tries\n",userGuess)
			break
		}

		

		

		
	}
}