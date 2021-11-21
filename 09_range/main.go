package main

import "fmt"

func main() {
	// Slices
	ids := []int{33,76,54,23,11,2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d \n", i, id)
	}

	// Not using indexes
	for _, id := range ids {
		fmt.Printf("ID: %d \n", id)
	}

	// Range with map
	countries := map[string]string{"BR":"Brasil", "PT":"Portugal", "EUA":"Estados Unidos"}

	for k, v := range countries {
		fmt.Printf("%s: %s \n", k, v)
	}

	for k := range countries {
		fmt.Printf("Siglas: %s \n", k)
	}
}
