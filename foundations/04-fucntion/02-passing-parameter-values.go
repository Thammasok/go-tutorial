package main

import (
	"fmt"
)

// pass-by-reference
func half(val *float64) {
	fmt.Printf("call half(%f)\n", *val)
	*val = *val / 2
}

// Anonymous Functions
var (
	mul = func(op0, op1 int) int {
		return op0 * op1
	}

	sqr = func(val int) int {
		return mul(val, val)
	}
)

// func main() {
//	// 1. pass-by-reference
// 	num := 2.807770
// 	fmt.Printf("num=%f\n", num)
// 	half(&num)
// 	fmt.Printf("half(num)=%f\n", num)

// 	// 2. Anonymous
// 	fmt.Printf("mul(25,7) = %d\n", mul(25, 7))
// 	fmt.Printf("sqr(13) = %d\n", sqr(13))
// 	fmt.Printf(
// 		"94 (°F) = %.2f (°C)\n",
// 		func(f float64) float64 {
// 			return (f - 32.0) * (5.0 / 9.0)
// 		}(94), // Send value 94
// 	)

// 	// 3. Closures ใช้งานภายใน Code Block นั้นๆ
// 	for i := 0.0; i < 360.0; i += 45.0 {
// 		rad := func() float64 {
// 			return i * math.Pi / 180
// 		}()
// 		fmt.Printf("%.2f Deg = %.2f Rad\n", i, rad)
// 	}
// }
