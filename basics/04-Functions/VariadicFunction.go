package main

import (
	"fmt"
	"reflect"
)

// Passing multiple atributes using a variadic function
func printMultipleStrings(strs ...string) {
	for _, str := range strs {
		fmt.Println(str)
	}
}

// Pass multiple different datatypes
func printMultipleVariables(vars ...interface{}) {
	for _, v := range vars {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

func main() {
	printMultipleStrings("Hello", "World")
	printMultipleStrings("Hello", "World", "!")

	printMultipleVariables(1, "Hello", true)
	printMultipleVariables(1, "Hello", true, 3.14)
}

