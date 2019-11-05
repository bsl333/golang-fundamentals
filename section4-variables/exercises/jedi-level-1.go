package main

import "fmt"

var x2 int = 42
var y2 string = "James Bond"
var z2 bool = true

type mytype int

var x4 mytype

var y5 int

func main() {
	/****
	* Exercise 1
	*****/
	fmt.Println("EX 1")
	x := 42
	y := `James Bond`
	z := true
	fmt.Println("\t", x, y, z)
	fmt.Println("\t", x)
	fmt.Println("\t", y)
	fmt.Println("\t", z)
	fmt.Printf("\v")

	/****
	* Exercise 2 & 3
	*****/
	fmt.Println("EX 2 & 3")
	s := fmt.Sprintf("\t%v\t%v\t%v", x2, y2, z2)
	fmt.Println(s)
	fmt.Printf("\v")

	/****
	* Exercise 4
	*****/
	fmt.Println("EX 4")
	fmt.Printf("\t%v\n\t%T\n", x4, x4)
	x4 = 42
	fmt.Println("\t", x)
	fmt.Printf("\v")

	/****
	* Exercise 5
	*****/
	fmt.Println("EX 5")
	y5 = int(x)
	fmt.Printf("\t%v\n\t%T\n", y5, y5)
	fmt.Printf("\v")

	/****
	* Exercise 6
	*****/

}
