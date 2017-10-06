package main

import "fmt"

func reverse{
	data := []rune(value)
	result := []rune{}

	//for loop that adds runes in reverse order
	for i := len(data) - 1; i >= 0; i--{
		result = append(result,data[i])
	}

	return string (result)

}