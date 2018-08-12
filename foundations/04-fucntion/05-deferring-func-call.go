package main

import "fmt"

// defer เลื่อนขั้นตอนการทำงาน
/*
* Such as:
* Closing open files
* Releasing network resources
* Closing the Go channel
* Committing database transactions
* And do on
 */

// func load(fname string) ([]string, error) {
// 	...
// 		file, err := os.Open(fname)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer file.Close()
// 	...
// }

func do(steps ...string) {
	defer fmt.Println("All done!")
	for _, s := range steps {
		defer fmt.Println(s)
	}

	fmt.Println("Starting")
}

func main() {
	do(
		"Find key",
		"Aplly break",
		"Put key in ignition",
		"Start car",
	)
}
