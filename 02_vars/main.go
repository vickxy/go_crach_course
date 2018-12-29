package main

import "fmt"

func main() {
	// Main Type
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128
	var age int32 = 25
	const isCool = true

	// Shorthand
	name, email := "Vikesh", "yadav.vikesh27@gmail.com"
	size := 32.12

	fmt.Println(name, age, email)
	fmt.Printf("%T\n", size)

}