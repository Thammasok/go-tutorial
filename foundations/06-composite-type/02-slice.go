package main

// Copying slices
func clone(v []float64) (result []float64) {
	result = make([]float64, len(v), cap(v))
	copy(result, v)
	return result
}

// func main() {
// 	// var months []string
// 	var list []int = []int{1, 2, 3, 4, 5}

// 	halfyr := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}

// 	// Slicing
// 	// <slice or array value>[<low_index>:<high_index>]
// 	// Condition >= && <
// 	fmt.Println(list[1:3])
// 	fmt.Println(halfyr[1:])

// 	// <slice_or_array_value>[<low_index>:<high_index>:max]
// 	var (
// 		months [12]string = [12]string{
// 			"Jan", "Feb", "Mar", "Apr", "May", "Jun",
// 			"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
// 		}
// 		summer = months[6:9:10]
// 	)

// 	// The max attribute specifies the index value to be used as the maximum capacity of the new slice.
// 	fmt.Println(summer)
// 	fmt.Println("summer Len: ", len(summer))
// 	fmt.Println("summer Cap: ", cap(summer))

// 	//keyword make
// 	list = make([]int, 5)
// 	fmt.Println("Make: ", list[:])

// 	// Appending to slices
// 	halfyr = append(halfyr, "Jul", "Aug", "Sep")
// 	halfyr = append(halfyr, []string{"Oct", "Nov", "Dec"}...)
// 	fmt.Println("Len:", len(halfyr), "Cap:", cap(halfyr), "halfyr: ", halfyr)
// }
