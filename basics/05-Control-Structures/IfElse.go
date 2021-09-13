package main

import "fmt"

func main() {
	x := true
	if x {
		fmt.Println("x is true")
	}

	y := 100
	if y > 10 {
		fmt.Println("y is greater than 10")
	} else {
		fmt.Println("y is less than 10")
	}

	grade := 10
	if grade == 10 {
		fmt.Println("You got an A")
	} else if grade == 6 {
		fmt.Println("You got a B")
	} else {
		fmt.Println("You failed")
	}

	if a := 10; a == 10 {
		fmt.Println("a is 10")
	} else {
		fmt.Println("a is not 10")
	}
}