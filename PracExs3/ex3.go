package main

import (
	"fmt"
)

func main() {
	fmt.Println("Example 1:")
	ex1()
	fmt.Println("\n")
	fmt.Println("Example 2:")
	ex2()
	fmt.Println("\n")
	fmt.Println("Example 3:")
	ex3()
	fmt.Println("\n")
	fmt.Println("Example 4:")
	ex4()
	fmt.Println("\n")
	fmt.Println("Example 5:")
	ex5()
	fmt.Println("\n")
	fmt.Println("Example 6:")
	ex6()
	fmt.Println("\n")
	fmt.Println("Example 7:")
	ex7()
	fmt.Println("\n")
	fmt.Println("Example 8:")
	ex8()
	fmt.Println("\n")
	fmt.Println("Example 9:")
	ex9()
	fmt.Println("\n")
	fmt.Println("Example 10:")
	ex10()
	fmt.Println("\n")
}

func ex1() {
	for i := 65; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func ex2() {
	for i2 := 65; i2 <= 90; i2++ {
		fmt.Println(i2)
		for j2 := 0; j2 < 3; j2++ {
			fmt.Printf("\t%#U\n", i2)
		}
	}
}

func ex3() {
	i3 := 1997
	for i3 <= 2021 {
		fmt.Println(i3)
		i3++
	}

}

func ex4() {
	i4 := 1997
	for {
		if i4 > 2021 {
			break
		}
		fmt.Println(i4)
		i4++
	}
}

func ex5() {
	i5 := 10
	for i5 <= 100 {
		m5 := i5 % 4
		fmt.Println(m5)
		i5++
	}
}

func ex6() {
	i6 := 0
	for i6 <= 20 {

		mod3 := i6 % 3
		mod4 := i6 % 4
		mod5 := i6 % 5

		if mod3 == 0 {
			fmt.Printf("%v is divisible by 3 with no remainder\n", i6)
		}

		if mod4 == 0 {
			fmt.Printf("%v is divisible by 4 with no remainder\n", i6)
		}

		if mod5 == 0 {
			fmt.Printf("%v is divisible by 5 with no remainder\n", i6)
		}
		i6++
	}
}

func ex7() {
	i6 := 0
	for i6 <= 20 {

		mod3 := i6 % 3
		mod4 := i6 % 4

		if mod3 == 0 {
			fmt.Printf("%v is divisible by 3 with no remainder\n", i6)
		} else if mod4 == 0 {
			fmt.Printf("%v is divisible by 4 with no remainder\n", i6)
		} else {
			fmt.Printf("%v is not divisible by 3 or 4\n", i6)
		}
		i6++
	}
}

func ex8() {
	switch {
	case false:
		fmt.Println("This should not print")
	case true:
		fmt.Println("This should print")
	}
}

func ex9() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("Go to the mountains")
	case "swimming":
		fmt.Println("Go to the pool")
	case "surfing":
		fmt.Println("go to hawaii")
	case "wingsuit flying":
		fmt.Println("I will plan your funeral")
	}
}

func ex10() {
	fmt.Println(true && true)  // true
	fmt.Println(true && false) // false
	fmt.Println(true || true)  // true
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
}
