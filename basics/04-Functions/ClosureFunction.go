package main

import "fmt"

func main() {
	l := 10
	b := 10

	func() {
		var area int
		area = l * b
		fmt.Println("Area of rectangle is:", area)
	}()
}