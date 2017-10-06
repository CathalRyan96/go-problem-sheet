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

func isP(s string) string{
	mid := len(s) / 2
	last := len(s) - 1

	for i := 0; i< mid; i++ {
		if s[i] != s[last-i]{
		return "No, this is not a palidrome"			
		}
		
	}
		return "Yes, this is a palidrome"
}