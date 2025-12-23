package main

import (
	"fmt"
	"math"
)

// s is a constant string that cannot be changed after being declared.
const s string = "constant"

func main() {

	// Print the value of s to the console.
	fmt.Println(s)

	// n is a constant integer that cannot be changed after being declared.
	const n = 500000000

	// d is a constant floating point number that cannot be changed after being declared.
	// It is calculated as 3e20 divided by n.
	const d = 3e20 / n
	fmt.Println(d)

	// Print the value of d as an int64 to the console.
	fmt.Println(int64(d))

	// Print the sine of n to the console.
	fmt.Println(math.Sin(n))
}
