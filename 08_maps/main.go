package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign kv
	// emails["bob"] = "bob@gmail.com"
	// emails["vikesh"] = "vikesh@gmail.com"
	
	// Declare and assign kv
	emails := map[string]string{"bob":"bob@foo.com", "vikesh":"vikesh@foo.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["bob"])
	delete(emails, "bob")
	fmt.Println(emails)

}