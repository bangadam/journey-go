package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	// ranges
	strings := []string{"a", "b", "c", "d", "e"}

	for i, s := range strings {
		fmt.Println(i, s)
	}

	// Getting index
	for i := range strings {
		fmt.Println(i)
	}

	// Getting value
	for _, s := range strings {
		fmt.Println(s)
	}
}