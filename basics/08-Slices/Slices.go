package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create empty slices
	var a []int
	fmt.Println(reflect.ValueOf(a).Kind()) // slice

	// Creating a slice using the make functions
	b := make([]string, 5, 10)
	fmt.Printf("a \tLen: %v \tCap: %v\n", len(a), cap(b)) // []string

	// Initialize the slice with values using a slice literal
	c := []int{1, 2, 3, 4, 5}
	fmt.Printf("c \tLen: %v \tCap: %v\n", len(c), cap(c)) // []int
	fmt.Println(c)                                        // [1 2 3 4 5]

	// Createing a slice using the keword
	var d = new([5]int)[0:2]
	fmt.Printf("d \tLen: %v \tCap: %v\n", len(d), cap(d)) // [5]int
	fmt.Println(d)                                        // [0 0]

	// add item using the append function
	var e = make([]string, 0, 5)
	fmt.Println(e)                                          // []
	e = append(e, "a", "b", "c")
	fmt.Println(e)                                          // [a b c]

	// Acces the slices
	var f = []int{1, 2, 3, 4, 5}
	fmt.Println(f[0])                                       // 1
	fmt.Println(f[1:3])                                     // [2 3]
	fmt.Println(f[2:])                                      // [3 4 5]

	// Change item values
	var g = []int{1, 2, 3, 4, 5}
	fmt.Println(g)                                          // [1 2 3 4 5]
	g[0] = 10
	fmt.Println(g)                                          // [10 2 3 4 5]

	// Copy slice into another slices
	var h = []int{1, 2, 3, 4, 5}
	var i = []int{6, 7, 8, 9, 10}
	copy(h, i)
	fmt.Println(h)                                          // [6 7 8 9 10]

	// Append a slice to an exsinting one
	var j = []int{1, 2, 3, 4, 5}
	var k = []int{6, 7, 8, 9, 10}

	j = append(j, k...)
	fmt.Println(j)                                          // [1 2 3 4 5 6 7 8 9 10]
}