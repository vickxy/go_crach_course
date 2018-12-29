package main

import "fmt"

func main() {
	// Array should be of fixed length
	// Type should be mentioned

	// var fruitArr [2]string

	// Assign values'
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"
	
	// Decalre and assign same time
	// numArr := [2]int{1,2}

	// Slices
	numSlice := []int{1, 2, 3}
	// fmt.Println(fruitArr)
	fmt.Println(len(numSlice))
	fmt.Println(numSlice[1:2])
} 