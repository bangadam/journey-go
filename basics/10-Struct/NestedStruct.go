package main

import "fmt"

// Nested struct
type Configuration struct {
	Env string
	Proxy Proxy
}

type Proxy struct {
	Address string
	Port string
}

func main() {
	// Creating an instance of a nested struct
	c := &Configuration{
		Env: "dev",
		Proxy: Proxy{
			Address: "127.0.0.1",
			Port: "8080",
		},
	}
	
	// Accessing the nested struct
	fmt.Println(c)
	fmt.Println(c.Env)
	fmt.Println(c.Proxy.Address)
}