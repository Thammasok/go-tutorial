package main

// import (
// 	"fmt"
// 	"sync"
// )

// func factorial(number int, wg *sync.WaitGroup) {
// 	answer := 1
// 	for i := number; i > 0; i-- {
// 		answer *= i
// 		// fmt.Println(i)
// 	}
// 	fmt.Printf("Fectorial of %d is %d\n", number, answer)
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	// เราต้องส่ง reference ของ wg ไปด้วย เพื่อที่จะสั่ง Done
// 	go factorial(5, &wg)
// 	go factorial(2, &wg)

// 	wg.Wait()
// }
