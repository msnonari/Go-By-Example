package main

import "fmt"

// plus takes two integers and returns their sum.
func plus(a int, b int) int {
	return a + b
}

// plusPlus takes three integers and returns their sum.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	// Call plus with arguments 1 and 2, and print the result.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	// Call plusPlus with arguments 1, 2, and 3, and print the result.
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
