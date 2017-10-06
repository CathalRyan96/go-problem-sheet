package main

import "fmt"
import "strings"

func main () {

	var test strings
	fmt.Println("Enter String to see if it is a Palimdrome!:")
	fmt.Scanf("%s\n",test)
	test = strings.ToLower(test)
	fmt.Println(isP(test))

	
}