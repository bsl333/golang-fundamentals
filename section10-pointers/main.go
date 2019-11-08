package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	// use & to see a variables address
	// use * to dereference address
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", &a)
	b := &a
	*b = 100
	fmt.Println(a)

	x := 0
	fmt.Println("x before mutate", x)
	mutateValue(&x)
	fmt.Println("x after mutated", x)
}

func mutateValue(num *int) {
	fmt.Println("num start in mutate", *num)
	*num = 3
	fmt.Println("num end in mutate", *num)
}
