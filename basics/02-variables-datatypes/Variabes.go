package main

import "fmt"

func main() {
	var a int
	var b string
	
	a = 10
	b = "Hello"

	fmt.Println(a)
	fmt.Println(b)

	var k int = 10

	fmt.Println(k)

	j := 50
	fmt.Println(j)

	firstName, LastName := "John", "Doe"
	fmt.Println(firstName, LastName)

	var (
		firstName1 string = "John"
		lastName1 string = "Doe"
	)

	fmt.Println(firstName1, lastName1)
}