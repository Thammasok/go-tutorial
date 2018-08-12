package main

import "fmt"

func readyToGo() {
	var readyToGo bool = false
	if !readyToGo {
		fmt.Println("Come on")
	} else {
		fmt.Println("Let's go!")
	}
}
