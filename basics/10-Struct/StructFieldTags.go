package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Name string `json:"name"`
	Age  int   `json:"age"`
}

func main() {
	json_string := `{"name":"Fido","age":3}`

	rocky := new(Dog)
	json.Unmarshal([]byte(json_string), rocky)
	fmt.Println(rocky)

	spot := new(Dog)
	spot.Name = "Spot"
	spot.Age = 5
	jsonStr, _ := json.Marshal(spot)
	fmt.Println(string(jsonStr))
}