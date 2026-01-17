package main

import "fmt"

// intSeq returns a function that generates a sequence of integers starting at 1.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// main is the entry point of our program.
func main() {
	// intSeq returns a function that generates a sequence of integers starting at 1.
	nextInt := intSeq()

	fmt.Println(nextInt()) // prints 1
	fmt.Println(nextInt()) // prints 2
	fmt.Println(nextInt()) // prints 3

	// Create a new sequence.
	newInts := intSeq()
	fmt.Println(newInts()) // prints 1
}
