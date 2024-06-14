package main

import "fmt"

// Empty main, just runs function
func main() {
	runProgram()
}

// Function to run program
func runProgram() {
	fmt.Println("Conditional statements in Go")
	fmt.Println("----------------------------------------------------")

	// Declare variable
	x := 10

	// If statment
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// Switch statement
	switch x {
	case 10:
		fmt.Println("x is 10")
	case 20:
		fmt.Println("x is 20")
	default:
		fmt.Println("x is neither 10 nor 20")
	}
}
