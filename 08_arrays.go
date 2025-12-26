package main

import "fmt"

func main() {

	// Declare an array of 5 integers and print its value.
	var a [5]int
	fmt.Println("emp:", a)

	// Set the value of the last element of a to 100 and print its value.
	a[4] = 100
	fmt.Println("set:", a)

	// Print the value of the last element of a.
	fmt.Println("get:", a[4])

	// Print the length of a.
	fmt.Println("len:", len(a))

	// Declare and initialize an array of 5 integers using the short variable declaration syntax.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Declare and initialize an array of 5 integers using the array literal syntax.
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Declare and initialize an array of 5 integers using the array literal syntax with index assignments.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Declare a 2D array of 2x3 integers and initialize it with values using a nested loop.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Declare and initialize a 2D array of 2x3 integers using the short variable declaration syntax.
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
