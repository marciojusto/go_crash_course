package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func printHello() {
	fmt.Println("Hello")
}

func main() {
	fmt.Println(greeting("Marcio"))
	fmt.Println(getSum(3, 4))
	printHello()
}
