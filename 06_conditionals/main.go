package main

import "fmt"

func main() {
	x := 15
	y := 10

	if lessThan(x, y) {
		fmt.Printf("%d is less than %d", x, y)
	} else {
		fmt.Printf("%d is less than %d \n", y, x)
	}

	whatIsTheColor("blue")
}

func lessThan(x int, y int) bool {
	return x < y
}

func whatIsTheColor(color string) {
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not blue or red")
	}
}
