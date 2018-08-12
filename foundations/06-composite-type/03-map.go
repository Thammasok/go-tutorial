package main

// import (
// 	"fmt"
// )

// func main() {
// 	// map[<key_type>]<element_type>
// 	var (
// 		histogram map[string]int = map[string]int{
// 			"Jan": 100, "Feb": 445, "Mar": 514, "Apr": 233,
// 			"May": 321, "Jun": 644, "Jul": 113, "Aug": 734,
// 			"Sep": 553, "Oct": 344, "Nov": 831, "Dec": 312,
// 		}

// 		table = map[string][]int{
// 			"Men":   []int{32, 55, 12, 55, 42, 53},
// 			"Women": []int{44, 42, 23, 41, 65, 44},
// 		}
// 	)

// 	// Get Value
// 	fmt.Println("Value in Key Jan: ", histogram["Jan"])

// 	fmt.Println("Len: ", len(histogram), "Data: ", histogram)
// 	fmt.Println(table)

// 	// Delete
// 	delete(table, "Women")
// 	fmt.Println("Len: ", len(table), "Data: ", table)

// 	// Make
// 	hist := make(map[int]string)
// 	hist[1] = "Jan"
// 	fmt.Println("Show Value After Make: ", hist)
// }
