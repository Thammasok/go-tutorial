package main

import (
	"math"
	"unsafe"
)

var _ int8 = 12       // -128 - 127
var _ int16 = -400    // -32768 - 32767
var _ int32 = 12022   // -2147483648 - 2147483647
var _ int64 = 1 << 33 // -9223372036854775808 - 9223372036854775807
var _ int = 3 + 1415  // 32 or 64-bit integers

var _ uint8 = 18                 // 0 - 255
var _ uint16 = 44                // 0 - 65535
var _ uint32 = 133121            // 0 - 4294967295
var i uint64 = 23113233          // 0 - 18446744073709551615
var _ uint = 7542                // 32 or 64-bit integers
var _ byte = 255                 // unit8
var _ uintptr = unsafe.Sizeof(i) // Unsigned size

var _ float32 = 0.5772156649 // Signed 32-bit (IEEE-754 standard)
var _ float64 = math.Pi      // Signed 64-bit (IEEE-754 standard)

var _ complex64 = 3.5 + 2i // size float32
var _ complex128 = -5.0i   // size float64

// func main() {
// 	fmt.Println("all types declared!")
// }
