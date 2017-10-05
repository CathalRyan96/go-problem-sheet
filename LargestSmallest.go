//Author Cathal Ryan	
package main

import "fmt"


func main(){
	x := []int{  //an array of random numbers
		50,62,24,37,
		21,89,55,91,
		18,6,33,77,
		80,35,4,12,
	}

	smallest, biggest := x[0],x[0] 
	
	for _, v := range x { // range loop adopted from https://stackoverflow.com/questions/34259800/is-there-a-built-in-min-function-for-a-slice-of-int-arguments-or-a-variable-numb/34259943
		if v > biggest {  
			biggest = v
		}

		if v < smallest {
			smallest = v
		}
	}

	fmt.Println("The largest number is ", biggest) //Returns largest number in the array
	fmt.Println("The smallest number is ", smallest) //Returns the smallest number in the array
    
}