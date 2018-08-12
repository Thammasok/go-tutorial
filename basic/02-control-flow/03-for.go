package main

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	Curr{"KES", "Kenyan Shilling", "Kenya", 404},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
}

// func main() {
// 	// For Condition used to semantically replace while and do...while loop
// 	// for x < 10 {
// 	// 	...
// 	// }

// 	// Infinite Loop
// 	// for {
// 	// 	...
// 	// }

// 	// Traditional
// 	// for x:=0; x < 10; x++ {
// 	// 	...
// 	// }

// 	// For range
// 	// var values = []int{1,2 , 3, 4, 5}
// 	// for i, val := range values {
// 	// 	...
// 	// }
// 	// Loop over map (in For range)
// 	// for k, v := range map[K]V {
// 	// 	...
// 	// }

// 	i := 0
// 	for i < len(currencies) {
// 		fmt.Println(currencies[i])
// 		i++
// 	}
// }
