package main

import "fmt"

// Empty main, just runs function
func main() {
	runProgram()
}

// Function to run program
func runProgram() {
	fmt.Println("Arrays in Go")
	fmt.Println("----------------------------------------------------")

	// Declare array
	var array [3]int

	// Populate array
	array[0] = 0
	array[1] = 1
	array[2] = 2

	// Output array
	fmt.Println("Array:", array)

}
