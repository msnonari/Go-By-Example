package main

import "fmt"

// fact returns the factorial of n.
func fact(n int) int {
	// base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}

	// recursive case: n! = n * (n-1)!
	return n * fact(n-1)
}

func main() {
	// print the factorial of 7
	fmt.Println(fact(7))

	// declare a function type for a Fibonacci sequence
	var fib func(n int) int

	// assign a function to the variable
	fib = func(n int) int {
		// base case: Fibonacci of 0 and 1 is 0 and 1 respectively
		if n < 2 {
			return n
		}

		// recursive case: Fibonacci of n is the sum of the Fibonacci of n-1 and n-2
		return fib(n-1) + fib(n-2)
	}

	// print the 7th Fibonacci number
	fmt.Println(fib(7))
}
