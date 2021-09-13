package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string to float
	s := "3.14"
	f, _ := strconv.ParseFloat(s, 8)
	fmt.Println(f)

	// String to boolean
	str := "true"
	b, _ := strconv.ParseBool(str)
	fmt.Println(b)

	// Float to string
	f = 3.14
	s = strconv.FormatFloat(f, 'f', -1, 64)
	fmt.Println(s)

	// Boolean to string
	b = true
	str = strconv.FormatBool(b)
	fmt.Println(str)

	// int to string
	i := 123
	str = strconv.Itoa(i)
	fmt.Println(str)

	// string to int
	str = "123"
	i, _ = strconv.Atoi(str)
	fmt.Println(i)

	// float to int
	f = 123.45
	i = int(f)
	fmt.Println(i)
}