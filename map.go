package main

import "fmt"

func main() {
	// Create a map
	m := make(map[string]int)

	// Insert key-value pairs
	m["apple"] = 5
	m["banana"] = 10
	m["orange"] = 15

	// Print the map
	fmt.Println("Map:", m)

	// Retrieve a value
	value, ok := m["apple"]
	if ok {
		fmt.Println("Value of 'apple':", value)
	} else {
		fmt.Println("Key 'apple' not found")
	}

	// Update a value
	m["banana"] = 20
	fmt.Println("Updated map:", m)

	// Delete a key-value pair
	delete(m, "orange")
	fmt.Println("Map after deletion:", m)

	// Check if a key exists
	if _, ok := m["grape"]; !ok {
		fmt.Println("Key 'grape' not found")
	}

	// Iterate over the map
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}
