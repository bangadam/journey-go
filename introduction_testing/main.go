package main

import "fmt"

func Caculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println(Caculate(3))
}