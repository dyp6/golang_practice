package main

import (
	"fmt"
)

func main() {
	fmt.Println("\nExample 1:")
	ex1()
	fmt.Println("\nExample 2:")
	ex2()
	fmt.Println("\nExample 3:")
	ex3()
	fmt.Println("\nExample 4:")
	ex4()
	fmt.Println("\nExample 5:")
	ex5()
}

// Example 1
func ex1() {
	x, y, z := 42, "James Bond", true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

// Example 2
// Automatically assigns value to variables called the zero-value
var x2 int
var y2 string
var z2 bool

func ex2() {
	fmt.Println(x2)
	fmt.Println("(", y2, ")")
	fmt.Println(z2)
}

var x3 int = 42
var y3 string = "James Bond"
var z3 bool = true

// Example 3
func ex3() {
	s := fmt.Sprintf("%v\t%v\t%v", x3, y3, z3)
	fmt.Println(s)
}

// Example 4
type funtype int

var x4 funtype

func ex4() {
	fmt.Println(x4)
	fmt.Printf("%T\n", x4)

	x4 = 42
	fmt.Println(x4)
}

// Example 5

var y5 int

func ex5() {
	y5 = int(x4)
	fmt.Println(y5)
	fmt.Printf("%T\n", y5)
}
