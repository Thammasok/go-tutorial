package main

import (
	"fmt"
)

/* การค้นหาไฟล์ต้องค้นหาไปทีละโฟลเดอร์ (เช่น Document, Image, Library) ตามลำดับ
แต่เราพบว่าเราสามารถ แบ่งการทำงานของโปรแกรมออกเป็น 3 ส่วนได้ คือ
1. ค้นหา Document
2. ค้นหา Image
3. ค้นหา Library อย่างเป็นอิสระต่อกัน
เรียกการโปรแกรมเพื่อแยกงานให้สามารถทำงานได้อย่างเป็นอิสระจากกันนี้ว่า Concurrency
ทำงานพร้อมกันเรียกว่า Parallelism
เราสามารถสร้างการทำงานแบบ Concurrency ได้ด้วยการใช้ Goroutines เพียงแค่เติม go
ตัวอย่าง

	func searchFromFolder(keyword string, folder string) {
		// ทำการค้นหา
	}

	func search(keyword string) {
		folders := [3]string{"Document", "Image", "Library"}

		// ค้นหาจากทีละโฟลเดอร์ตามลำดับ
		for _, folder := range folders {
			go searchFromFolder(keyword, folder)
		}
	}
*/

// 1. Gorotines
func count(start, stop, delta int) {
	// fmt.Println("---------")
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}

func main() {
	// count(10, 50, 10)
	// count(60, 100, 10)
	// count(110, 200, 20)

	// 2. Gorotines
	// go <function or expression>
	// go count(10, 50, 10)
	// fmt.Println("---------")
	// go count(60, 100, 10)
	// fmt.Println("---------")
	// go count(110, 200, 20)
	// fmt.Println("---------")
	// fmt.Scanln()
}
