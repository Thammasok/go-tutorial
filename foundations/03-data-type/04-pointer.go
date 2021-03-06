package main

import "fmt"

// 1. Type as a pointer
// var valPtr *float32
// var countPtr *int
// var person *struct {
// 	name string
// 	age  int
// }
// var matrix *[1024]int
// var row []*int64

// 2. The address operator
func addressOperator() {
	// aptr -> address of a
	var a int = 1024
	var aptr *int = &a

	fmt.Printf("a=%v\n", a)
	fmt.Printf("aptr=%v\n", aptr)
}

// 3. The new() function
func newFunction() {
	intptr := new(int)
	*intptr = 44

	p := new(struct{ first, last string })
	p.first = "Samuel"
	p.last = "Pierre"

	fmt.Printf("Value %d, type %T\n", *intptr, intptr)
	fmt.Printf("Person %+v\n", p)
}

func main() {
	addressOperator()
	newFunction()
}
