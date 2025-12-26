package main

import (
	"fmt"
	"time"
)

// main is the entry point of our program.
func main() {

	// Print the number 2 as a string.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Print whether it is the weekend or a weekday.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Print whether it is before or after noon.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// whatAmI is a function that takes an interface argument and prints its type.
	// It uses a type switch to determine the type of the argument.
	whatAmI := func(i interface{}) {
		// Use a type switch to determine the type of the argument.
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	// Call whatAmI with different types of arguments.
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
