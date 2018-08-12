package main

// func <func-identifier>(<argument-list>) [<result-list>] { [return] [<value or expression list] }
// Basic
func twoSum(numberOne, numberTwo int) int {
	return numberOne + numberTwo
}

// Variadic
func avg(nums ...float64) float64 {
	n := len(nums)
	t := 0.0
	for _, v := range nums {
		t += v
	}
	return t / float64(n)
}

// Function result parameters
func div(op0, op1 int) (int, int) {
	r := op0
	q := 0
	for r >= op1 {
		q++
		r = r - op1
	}
	return q, r
}

// func main() {
// 	fmt.Printf("%d + %d = %d\n", 5, 1, twoSum(5, 1))

// 	//Variable length arguments
// 	fmt.Printf("avg([1, 2.5, 3.75]) =%.2f\n", avg(1, 2.5, 3.75))

// 	// Function result parameters
// 	q, r := div(71, 5)
// 	fmt.Printf("div(71,5) -> q = %d, r = %d\n", q, r)
// }
