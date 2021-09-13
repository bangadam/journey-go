package main

import "fmt"

var m = map[string]string{
	"foo": "bar",
	"baz": "qux",
}

func main() {
	// Declaring empty map
	var shoppingList = map[string]int{}
	fmt.Println(shoppingList)

	// Initializing a map
	var people = map[string]int{"Elon": 46, "Bill": 42}
	fmt.Println(people)

	// Map declaration using make function
	var peopleList = make(map[string]int)
	peopleList["Elon"] = 46
	peopleList["Bill"] = 42
	fmt.Println(peopleList)

	// Accessing Items map
	fmt.Println(m["foo"])
	fmt.Println(m["baz"])

	// adding items
	m["foo2"] = "bar2"
	fmt.Println(m)

	// Deleting items map
	delete(m, "foo")
	fmt.Println(m)

	// Iterating over map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Test if an item exists
	c, ok := m["baz"]
	fmt.Println(c, ok)
}