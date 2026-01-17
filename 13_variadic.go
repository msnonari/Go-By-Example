package main

import "fmt"

// sum takes a variable number of integers as arguments and returns their sum.
func sum(nums ...int) {
	// Print the arguments passed to the function.
	fmt.Print(nums, " ")

	// Initialize a variable to store the total sum.
	total := 0

	// Iterate over the arguments and add them to the total sum.
	for _, num := range nums {
		total += num
	}
	// Print the total sum.
	fmt.Println(total)
}

func main() {
	// Call sum with two arguments.
	sum(1, 2)

	// Call sum with three arguments.
	sum(1, 2, 3)

	// Create a slice of integers.
	nums := []int{1, 2, 3, 4}

	// Call sum with the slice of integers as an argument using the spread operator (...).
	sum(nums...)
}
