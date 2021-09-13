package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func getArea(x int, y int) (area int) {
	area = x * y
	return
}

func rectangleArea(x int, y int) (area int, perimeter int) {
	area = x * y
	perimeter = (x + y) * 2
	return
}

func addValue(x *int, y *string) {
	*x = *x + 1
	*y = *y + "World!"
	return
}

func main() {
	sum := add(10, 5)
	fmt.Println(sum)
	
	area := getArea(10, 5)
	fmt.Println(area)

	area, perimeter := rectangleArea(10, 5)
	fmt.Println(area, perimeter)

	var number = 30
	var halo = "Hello"
	addValue(&number, &halo)
	fmt.Println(number, halo)
}