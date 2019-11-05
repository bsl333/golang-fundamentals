package main

import (
	"fmt"
	"runtime"
)

// to declare mulitple consts

const (
	pi  = 3.141529
	tau = 2 * pi
)

// iota example with bit shifting

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("kb = %b\n", kb)
	fmt.Printf("mb = %b\n", mb)
	fmt.Printf("gb = %b\n", gb)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOARCH)

	a := "az"

	aBytes := []byte(a)

	for i := 0; i < len(aBytes); i++ {
		fmt.Printf("%x\n", aBytes[i])
	}
	fmt.Printf("%X\n", 350)
	greet("Good moring", "Rex")

	fmt.Println(areaCircle(3), circumferenceCircle(3))
}

func greet(greeting, thing string) {
	fmt.Println(greeting + ", " + thing)
}

func areaCircle(radius float64) float64 {
	return pi * radius * radius
}

func circumferenceCircle(radius float64) float64 {
	return tau * radius
}
