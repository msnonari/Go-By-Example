package main

import (
	"fmt"
	"maps"
)

// main is the entry point of our program.
func main() {
	// Create a map of strings to integers.
	m := make(map[string]int)

	// Set the value of the key "k1" to 7 and the value of the key "k2" to 13.
	m["k1"] = 7
	m["k2"] = 13

	// Print the value of the map.
	fmt.Println("map:", m)

	// Retrieve the value of the key "k1" from the map.
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Retrieve the value of the key "k3" from the map. If the key does not exist, the value will be zero.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Print the length of the map, which is the number of key-value pairs.
	fmt.Println("len:", len(m))

	// Delete the key-value pair with the key "k2" from the map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// Clear the map of all key-value pairs.
	clear(m)
	fmt.Println("map:", m)

	// Retrieve the value of the key "k2" from the map. If the key does not exist, the value will be zero.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Create a new map with the key-value pairs {"foo": 1, "bar": 2}.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Create a new map with the same key-value pairs as n.
	n2 := map[string]int{"foo": 1, "bar": 2}
	// Check if the two maps are equal.
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
