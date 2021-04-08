package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercise 1:")
	ex1()
	fmt.Println("\n")
	fmt.Println("Exercise 2:")
	ex2()
	fmt.Println("\n")
	fmt.Println("Exercise 3:")
	ex3()
	fmt.Println("\n")
	fmt.Println("Exercise 4:")
	ex4()
	fmt.Println("\n")
}

var x1 int

func ex1() {
	x1 = 400
	d := fmt.Sprintf("Decimal: %d", x1)
	b := fmt.Sprintf("Binary: %b", x1)
	h := fmt.Sprintf("Hex: %#x", x1)
	fmt.Println(d, "\t", b, "\t", h, "\t")
}

func ex2() {
	a2 := (42 == 42)
	b2 := (42 <= 43)
	c2 := (42 >= 43)
	d2 := (42 != 42)
	e2 := (42 > 41)
	f2 := (42 < 41)
	fmt.Println(a2, "\t", b2, "\t", c2, "\t", d2, "\t", e2, "\t", f2)
}

const a3 = "bar"
const b3 string = "bar"
const c3 = 42
const d3 int = 42

func ex3() {
	fmt.Println(a3, "\t", b3, "\n", c3, "\t", d3)
}

func ex4() {
	fmt.Println(" ")
}
