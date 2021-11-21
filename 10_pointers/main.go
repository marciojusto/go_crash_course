package main

import "fmt"

func main() {
	// Use & to get the memory address of the attribute
	a := 5
	b := &a

	fmt.Println(a, b)

	// Use * to get the value from the memory address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 50
	fmt.Println(a)
}
