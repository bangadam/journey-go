package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}

	// Copy the entire array to y
	y := x

	// Copy by reference
	z := &x

	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)

	x[0] = 1

	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("z: %v\n", *z)
}