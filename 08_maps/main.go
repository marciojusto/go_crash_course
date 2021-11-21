package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign key/value
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// Declare map and add key/value
	countries := map[string]string{"BR":"Brasil", "PT":"Portugal", "EUA":"Estados Unidos"}
	countries["CH"] = "China"

	fmt.Println(countries)
}
