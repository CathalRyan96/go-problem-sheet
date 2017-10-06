package main

import "fmt"
import "strings"

//main function that takes in users string
func main () {

	var ip string
	fmt.Println("Enter String to see if it is a Palimdrome!:")
	fmt.Scanf("%s\n", &ip)
	ip = strings.ToLower(ip)
	fmt.Println(isP(ip))

	
}

//Function to delcare whether the string is a palidrome 
func isP(s string) string{
	mid := len(s) / 2  //initialized a mid variable that is half the he length of the string
	last := len(s) - 1 // initialized a last variable that is the length of the string - 1

	for i := 0; i< mid; i++ {
		if s[i] != s[last-i]{
		return "No, this is not a palidrome"			
		}
		
	}
		return "Yes, this is a palidrome"
}