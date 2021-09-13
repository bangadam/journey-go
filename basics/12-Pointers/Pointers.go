package main

import "fmt"

func main() {
	i, j := 42, 2701

	// Point to i
	p := &i
	fmt.Println(*p)

	// Set i using the pointer function
	*p = 21
	fmt.Println(i)

	// Point to j
	p = &j
	*p = *p/7
	fmt.Println(j)
}