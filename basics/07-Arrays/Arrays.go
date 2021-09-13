package main

import "fmt"

func main() {
	// Declaring array
	var a [5]int

	// assigning array
	a[0] = 1
	a[1] = 2
	a[2] = 3

	// Accessing array
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])

	// Initializing array
	b := [5]int{1, 2, 3, 4, 5}
	var y [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(b)
	fmt.Println(y)

	// Initializing array with elipsis
	k := [...]int{1, 2, 3, 4, 5}

	fmt.Println(k)

	// Initializing array for spesific array element
	c := [5]int{1: 1, 2: 2, 3: 3, 4: 4}
	fmt.Println(c)
}