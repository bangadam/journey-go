package main

func multiply(x, y int) int {
	return x * y
}

// Function that returns another function
func partialMultiplication(x int) func(int) int {
	return func(y int) int {
		return multiply(x, y)
	}
}

func main() {
	// Call partialMultiplication with x = 2
	// and then call the returned function with y = 3
	// and print the result
	f := partialMultiplication(2)
	println(f(3))
}