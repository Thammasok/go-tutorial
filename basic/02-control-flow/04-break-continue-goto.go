package main

import (
	"fmt"
)

func main() {
	var a string
Start:
	for {
		switch {
		case a < "aaa":
			goto A
		case a >= "aaa" && a < "aaabbb":
			goto B
		case a == "aaabbb":
			// fmt.Println("Before break: ", a)
			break Start
		}
	A:
		a += "a"
		continue Start
	B:
		a += "b"
		continue Start
	}
	fmt.Println(a)
}
