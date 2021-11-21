package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

type Person2 struct {
	firstName, lastName, city, gender string
	age int
}

// greet method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasAge method (pointer receiver)
func (p *Person) hasAge() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Female" {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "John", lastName: "Smith", city: "Boston", gender: "Male", age: 25}
	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	// Alternative
	person2 := Person{"Samantha", "Smith", "Boston", "Female", 30}
	fmt.Println(person2)
	person2.hasAge()
	person2.getMarried("Williams")
	fmt.Println(person2.greet())
}
