package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	/*firstName string
	lastName string
	city string
	gender string 
	age int
	*/
	// Inshort
	firstName, lastName, city, gender string
	age int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " year old."
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age ++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "male"{
		return
	}else{
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	// person1 := Person{firstName: "vickxy", lastName: "yadav", city: "Bangalore", gender: "male", age: 25}
	
	// Alternative
	person1 := Person{"Dyna", "Robart", "Bangalore", "female", 25}	
	person1.age++
	fmt.Println(person1)

	// get single field
	fmt.Println(person1.firstName)

	person1.hasBirthday()
	fmt.Println(person1.greet())
	
	person1.getMarried("Williams")
	fmt.Println(person1.greet())
}