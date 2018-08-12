package main

import "fmt"

// ประกาศตัวแปร
// var <identifier list> <type>
func variableDeclaration() {
	fmt.Println("> > > > > > > > > > Variable declaration < < < < < < < < < <")

	var name, desc string
	var radius int32
	var mass float64
	var active bool
	var satellites []string
	// var number []int

	fmt.Println("Zero-Value ---------------------")
	fmt.Println("Name (string): ", name)
	fmt.Println("DESC (string): ", desc)
	fmt.Println("Radius (Int): ", radius)
	fmt.Println("Active (Bool): ", active)
	fmt.Println("Mass (float64): ", mass)
	fmt.Println("Satellites (Array Srting): ", satellites)
	// fmt.Println("Number (Array int): ", number)

	fmt.Println("Set Value ----------------------")

	name = "Sun"
	desc = "Star"
	radius = 685800
	mass = 1.989E+30
	active = true
	satellites = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	// number = []int{1, 2, 3, 4, 5}

	fmt.Println("Name: ", name)
	fmt.Println("Desc: ", desc)
	fmt.Println("Radius (km): ", radius)
	fmt.Println("Active: ", active)
	fmt.Println("Mass (kg): ", mass)
	fmt.Println("Satellites: ", satellites)
	// fmt.Println("Number: ", number)
}

// ประกาศตัวแปร พร้อมระบุค่า
// var <identifier list> <type> = <value list or initializer expressions>
func initializedDeclaration() {
	fmt.Println("> > > > > > > > > > Initialized declaration < < < < < < < < < <")

	// ประกาศตัวแปร พร้อมระบุค่าและระบุประเภท
	var name, desc string = "Earth", "Planet"
	var radius int32 = 6378
	var mass float32 = 5.972E+24
	var actives bool = true
	var satellites = []string{
		"Moon",
	}

	// fmt.Println("> > > > > > > > > > Omitting variable types < < < < < < < < < <")

	// ประกาศตัวแปร พร้อมระบุค่า แต่ไม่ระบุ type
	// var name, desc = "Mars", "Planet"
	// var radius = 6755
	// var mass = 641693000000000.0
	// var active = true
	// var satellites = []string{
	//   "Phobos",
	//   "Deimos",
	// }
	// OR //
	// var (
	//   name string = "Earth"
	//   desc string = "Planet"
	//   radius int32 = 6378
	//   mass float64 = 5.972E+24
	//   active bool = true
	//   satellites []string
	// )

	fmt.Println("Name: ", name)
	fmt.Println("Desc: ", desc)
	fmt.Println("Radius (km): ", radius)
	fmt.Println("Actives: ", actives)
	fmt.Println("Mass (kg): ", mass)
	fmt.Println("Satellites: ", satellites)
}

//ประกาศตัวแปรแบบสั้น
func shortVariableDeclaration() {
	fmt.Println("> > > > > > > > > > Short variable declaration < < < < < < < < < <")

	name := "Neptune"
	desc := "Planet"
	radius := 24764
	mass := 1.024e26
	active := true
	satellites := []string{
		"Naiad", "Thalassa", "Despina", "Galatea", "Larissa",
		"S/2004 N 1", "Proteus", "Triton", "Nereid", "Halimede",
		"Sao", "Laomedeia", "Neso", "Psamathe",
	}

	fmt.Println("Name: ", name)
	fmt.Println("Desc: ", desc)
	fmt.Println("Radius (km): ", radius)
	fmt.Println("Active: ", active)
	fmt.Println("Mass (kg): ", mass)
	fmt.Println("Satellites: ", satellites)
}

// func main() {
// 	variableDeclaration()
// 	initializedDeclaration()
// 	shortVariableDeclaration()
// }
