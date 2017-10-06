package main

import "fmt"

func reverse(value string) string{
	data := []rune(value)
	result := []rune{}

	//for loop that adds runes in reverse order
	for i := len(data) - 1; i >= 0; i--{
		result = append(result,data[i])
	}

	return string (result)

}

func main (){
value1 := "Cathal"  // setting string to Cathal
reverseString1 := reverse(value1) // reverses string
fmt.Printf("String is %s\n", value1) // prints out normal string

fmt.Printf("The reversed string is: %s\n",reverseString1) // prints out the reversed string




}
