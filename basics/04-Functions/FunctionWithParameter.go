package main

import "fmt"

func main() {
	hello("go")
	add(20,30)
}

func hello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func add(a int, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}