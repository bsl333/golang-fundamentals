package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("i is", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\tj is %v\n", j)
		}
		if i == 4 {
			fmt.Printf("\v\v")
		}
	}

	x := 1

	for x < 10 {
		fmt.Println(x)
		x += 5
	}

	arr := [3]int{1, 2, 4}
	for i, val := range arr {
		if val%2 == 1 {
			continue
		}
		fmt.Println(i, val)
	}

	for i := 33; i <= 122; i++ {
		// fmt.Printf("%v\t%c\n", i, i)
		// or...
		fmt.Printf("%v\t%#U\n", i, i)
	}

	// initialzation statement - x's scope is limited to the if block
	if x := 42; x > 10 {
		fmt.Println(x)
	}

	// if else

	someNum := 10
	if someNum > 15 {
		fmt.Println("someNum is greater than 15", someNum)
	} else if someNum > 10 {
		fmt.Println("someNum is greater than 10", someNum)
	} else {
		fmt.Println("someNum is less than or equal to 10", someNum)
	}

	// switch statements
	// fallthrough is not true by default, so no need to break in switch statements.
	// if using fallthrough, only falls through to next case, then breaks
	switch someNum {
	case 10, 30: // can check for multiple values in one case statement.
		fmt.Println("10 | 30")
		fallthrough
	case 15:
		fmt.Println("15")
	case 20:
		fmt.Println("20")
	default:
		fmt.Println("some other value")
	}

	switch {
	case someNum == 10:
		fmt.Println("10")
	case someNum == 15:
		fmt.Println("15")
	case someNum == 20:
		fmt.Println("20")
	default:
		fmt.Println("some other value")
	}
}
