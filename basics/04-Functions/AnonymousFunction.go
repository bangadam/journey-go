package main

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	area := area(100, 20)
	fmt.Println(area)
}