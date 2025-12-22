package main

import "fmt"

func main() {
	// Declare a variable with an initial value.
	var a = "initial"
	fmt.Println(a)

	// Declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Declare a variable with type inference.
	var d = true
	fmt.Println(d)

	// Declare a variable without an initial value.
	var e int
	fmt.Println(e)

	// Declare a short variable declaration.
	f := "apple"
	fmt.Println(f)
}
