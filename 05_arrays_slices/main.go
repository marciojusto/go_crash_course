package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign Values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	cars := []string{"Toyota", "Nissan", "BMW", "Peugeot"}

	fmt.Println(fruitArr)
	fmt.Println(cars)
	fmt.Println(cars[1:3]) // -> slice
}
