package main

// import [package name identifier] "<import path>"
import (
	"fmt"
	// _ "os"

	f "./foo"
	. "./foo/bar" //Should not use dot imports
	"./foo/bazz"
)

// Package initialization
func init() {
	fmt.Println("Main File.")
}

func main() {
	fmt.Println("Import package")

	f.PrintQux()
	bazz.PrintBee()
	PrintMe()
}
