package main

import "fmt"

func first() {
	fmt.Println("first function")
}

func Second() {
	fmt.Println("second function")
}

// Defer is a spezial statement that schedules functions to be executed after the function completes
func main() {
	defer Second()
	first()
}