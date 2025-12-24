package main

import "fmt"

func main() {

	// Check if 7 is even.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Check if 8 is divisible by 4.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Check if either 8 or 7 is even.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// Check the properties of the number 9.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
