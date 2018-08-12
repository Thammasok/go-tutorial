package main

import "fmt"

var greetings = [][]string{
	{"Hello, World!!!", "English"},
	{"สวัสดีโลก!!!", "Thai"},
	{"Salut Monde", "French"},
}

func greeting(language string) []string {
	for i := range greetings {
		if greetings[i][1] == language {
			return greetings[i]
		}
	}

	return greetings[0]
}

func main() {
	g := greeting("Thai")
	fmt.Printf("%s (%s)\n", g[0], g[1])
}

/***************** The Package Clause *********************/
//1. Package Clause
//package main

//2. Import Declaration
// import "fmt"
// or
// import (
// 	"fmt"
// )

//3. Source Body
// var greetings = [][]string{
// 	{"Hello, World!","English"},
// 	...
// }

// func greeting() [] string {
// 	...
// }

// func main() {
// 	...
// }

/**************** Optional semicolon & Multiple lines ************/
// lonStr := "Hello World! " +
// "How are you?"

//  fmt.Printf("[%s] %d %d %v",
//  str,
//  num1,
//  num2,
//  nameMap)

// 	weekDays2 := []string{
// 		"Mon", "Tue",
// 		"Wed", "Thr",
// 		"Fri",
// 	}
