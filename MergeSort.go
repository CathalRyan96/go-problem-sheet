//Adopted from : https://gobyexample.com/sorting
package main

import "fmt"
import "sort"

func main() {
	s := make([]int, 0, 5) //creating two slices
	s2 := make([]int, 0, 5)

	s = append(s, 10, 8, 9) // filling the slices
	s2 = append(s2, 1, 3, 2)
	s3 := make([]int, 0, 10)

	s3 = append(s, s2[0], s2[1], s2[2])
	fmt.Print(s3) // prints out normal merged list

	sort.Ints(s3) 

	fmt.Println(s3) // prints out sorted list
	


}