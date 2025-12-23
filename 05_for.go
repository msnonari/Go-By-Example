package main

import "fmt"

func main() {

	// Initialize a variable and loop from 1 to 3.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Use a for loop with a range to loop from 0 to 3.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Use a range to loop from 0 to 3.
	for i := range 3 {
		fmt.Println("range", i)
	}

	// Use an infinite loop and break it after one iteration.
	for {
		fmt.Println("loop")
		break
	}

	// Use a range to loop from 0 to 6 and use a continue statement to skip even numbers.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
