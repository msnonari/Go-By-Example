package main

import (
	"fmt"
	"slices"
)

func main() {
	// Declare an uninitialized slice
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Create a slice with make, specifying length and capacity
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Assign values to slice elements
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// Check slice length
	fmt.Println("len:", len(s))

	// Append elements to the slice (may reallocate)
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Copy slice contents to a new slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Create slice from index 2 to 5 (exclusive)
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Create slice from start to index 5 (exclusive)
	l = s[:5]
	fmt.Println("sl2:", l)

	// Create slice from index 2 to end
	l = s[2:]
	fmt.Println("sl3:", l)

	// Declare and initialize a slice literal
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Compare two slices for equality
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Create a 2D slice (slice of slices)
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
