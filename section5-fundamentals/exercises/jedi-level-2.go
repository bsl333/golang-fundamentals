package main

import "fmt"

// for ex 3 - signed vs unsigned
const (
	untyped uint8 = 3
	typed   int8  = -3
)

// for ex 6 - iota
const baseYear = 2019
const (
	year1 = iota + baseYear
	year2
	year3
	year4
)

func main() {
	/*****
	* Exercise 1: print a number in decimal, binary, and hex
	*****/
	num := 42
	fmt.Println("Ex 1")
	types := [3]string{"decimal = %d", "binary = %b", "hex = %#x"}
	printMultipleFormats(types, num)

	/*****
	* Exercise 2: write expressions using equality, greater/less than checks
	*****/
	fmt.Println("Ex 2")

	x := 3
	y := 5
	isEqual := x == y
	isNotEqual := x != y
	isLessThanOrEqual := x <= y
	isGreaterThanOrEqual := x >= y
	isLessThan := x < y
	isGreaterThan := x > y

	fmt.Println("\t", isEqual, isNotEqual, isLessThanOrEqual, isGreaterThanOrEqual, isLessThan, isGreaterThan)

	/*****
	* Exercise 3: create typed and untyped constants
	*****/

	fmt.Println("Ex 3")
	fmt.Println("\ttyped:  ", typed)
	fmt.Println("\tuntyped: ", untyped)

	/*****
	* Exercise 4: bit shifting
	*****/

	fmt.Println("Ex 4")
	num = 1523
	printMultipleFormats(types, num)
	shifted := num << 1
	printMultipleFormats(types, shifted)

	/*****
	* Exercise 5: create a string using raw string literals
	*****/
	fmt.Println("Ex 5")
	name := `Branden Lowe`
	fmt.Println("\t", name)

	/*****
	* Exercise 6: iota - create consts for the next 4 years
	*****/

	fmt.Println("Ex 6")
	fmt.Println("\t", year1, year2, year3, year4)
}

func printMultipleFormats(types [3]string, num int) {
	for i := 0; i < len(types); i++ {
		fmt.Printf("\t"+types[i]+"\n", num)
	}
}
