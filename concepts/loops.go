package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")
	fmt.Println("----------------------------------------------------")
	// Equivalent to while loop
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// C++ style for loops
	for y := 0; y < 3; y++ {
		fmt.Println(y)
	}

}
