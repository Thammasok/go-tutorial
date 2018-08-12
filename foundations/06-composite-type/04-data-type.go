package main

import (
	"fmt"
)

// 2. The anonymous field
type age int

// 1. Basic
type person struct {
	name    string
	age     // anonymous
	address address
}

type address struct {
	street      string
	city, state string
	postal      string
}

func makePerson() person {
	addr := address{
		city:   "Bang na",
		state:  "Bangkok",
		postal: "10260",
	}
	return person{
		name:    "Thammasok",
		age:     25, // anonymous
		address: addr,
	}
}

func main() {
	// 1. Basic
	result := makePerson()
	fmt.Println(result)

	// 3. Structs as parameters
	result.name = "Kukkuk"
	fmt.Println("Result after update: ", result)
}

// 4. Field tags
// type Person struct {
// 	Name    string `json:"person_name"`
// 	Title   string `json:"person_title"`
// 	Address `json:"person_address_obj"`
// }
