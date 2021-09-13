package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Company string
}

func (conf Employee) EmployeeInfo() string {
	fmt.Println("Employee: ", conf.Name, conf.Age, conf.Company)
	return "------------------"
}

func main() {
	c := &Employee{
		Person: Person{
			Name: "John",
			Age:  30,
		},
		Company: "Google",
	}

	fmt.Println(c.EmployeeInfo())
}
