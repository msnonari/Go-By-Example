package main

import "fmt"

// It demonstrates the use of strings, numbers, and booleans.
func main() {

	// Concatenate strings
	fmt.Println("go" + "lang")

	// Perform arithmetic operations
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Use boolean operators
	fmt.Println(true && false) // AND operator
	fmt.Println(true || false) // OR operator
	fmt.Println(!true)         // NOT operator
}
