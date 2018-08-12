package main

import (
	"fmt"
)

func main() {
	number := [5]int{5, 7, 2, 5, 1}
	sum := 0
	for _, n := range number {
		sum += n
	}

	fmt.Println(sum)
}
