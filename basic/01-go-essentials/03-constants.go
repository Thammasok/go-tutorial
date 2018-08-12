package main

import (
	"math"
	"time"
)

func constantVariable() {
	// ประกาศค่าคงที่แบบระบุประเภท
	// const <identifier list> type = <value list or initializer expressions>
	const a1, a2 string = "Mastering", "Go"
	const b rune = 'G'
	const c bool = false
	const d int32 = 111009
	const e float32 = 2.71828
	const f float64 = math.Pi * 2.0e+3
	const g complex64 = 5.0i
	const h time.Duration = 4 * time.Second

	// ประกาศค่าคงที่แบบไม่ระบุประเภท
	// const <identifier list> = <value list or initializer expression>
	const i = "G is" + " for Go "
	const j = 'V'
	// const k1, k2 = true, !k1
	const l = 111*100000 + 9
	const m1 = math.Pi / 3.141592
	// const m2 = 1.414213562373095048801688724209698078569671875376...
	// const m3 = m2 * m2
	// const m4 = m3 * 1.0e+400
	const n = -5.0i * 3
	const o = time.Millisecond * 5

	// Assign Untyped const
	// const m2 = 1.414213562373095048801688724209698078569671875376...
	// var u1 float32 = m2
	// var u2 float64 = m2
	// u3 := m2

	// const (
	// 	a1, a2 string        = "Mastering", "Go"
	// 	b      rune          = 'G'
	// 	c      bool          = false
	// 	d      int32         = 111009
	// 	e      float32       = 2.71828
	// 	f      float64       = math.Pi * 2.0e+3
	// 	g      complex64     = 5.0i
	// 	h      time.Duration = 4 * time.Second
	// ...
	// )
}
