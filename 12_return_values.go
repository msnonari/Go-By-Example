package main

import "fmt"

// vals returns two integers.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Call vals and assign the returned values to a and b.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Call vals and ignore the first returned value.
	// Assign the second returned value to c.
	_, c := vals()
	fmt.Println(c)
}
