//Author Cathal Ryan
//Adopted from https://tour.golang.org/welcome/4
package main

import (
	"fmt" // imports fmt package which implements formatted I/O

	"time" // imports time package that provides funcionality for measuring and displaying time
	
)

func main() {
	fmt.Println("This program displays the current time!") 

	fmt.Println("The time is", time.Now()) // Prints the current time to the console
}