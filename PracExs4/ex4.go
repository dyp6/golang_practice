package main

import (
	"fmt"
)

func main() {
	fmt.Println("Example 1:")
	ex1()
	fmt.Println()
	fmt.Println("Example 2:")
	ex2()
	fmt.Println()
	fmt.Println("Example 3:")
	ex3()
	fmt.Println()
	fmt.Println("Example 4:")
	ex4()
	fmt.Println()
	fmt.Println("Example 5:")
	ex5()
	fmt.Println()
	fmt.Println("Example 6:")
	ex6()
	fmt.Println()
	fmt.Println("Example 7:")
	ex7()
	fmt.Println()
	fmt.Println("Example 8:")
	ex8()
	fmt.Println()
	fmt.Println("Example 9:")
	ex9()
	fmt.Println()
	fmt.Println("Example 10:")
	ex10()
	fmt.Println()
}

func ex1() {
	x1 := [5]int{41, 42, 43, 44, 45}
	for i, v := range x1 {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x1)
}

func ex2() {
	x2 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x2 {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x2)
}

func ex3() {
	x3 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x3[:5])
	fmt.Println(x3[5:])
	fmt.Println(x3[2:7])
	fmt.Println(x3[1:6])
}

func ex4() {

}

func ex5() {

}

func ex6() {

}

func ex7() {

}

func ex8() {

}

func ex9() {

}

func ex10() {

}
