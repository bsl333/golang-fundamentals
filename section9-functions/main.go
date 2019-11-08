package main

import (
	"fmt"
	"math"
)

type person struct {
	first, last string
}

type secretAgent struct {
	person
	ltk bool
}

// func (s secretAgent) speak(message string) {
// 	fmt.Println(message)
// 	fmt.Println("I am", s.first, s.last)
// }

func (p person) speak(message string) {
	fmt.Println(message)
	fmt.Println("I am", p.first, p.last)
}

type human interface {
	speak(string)
}

func printer(el interface{}) {
	fmt.Println(el)
}

func bar(h human) {
	fmt.Println("I am a human")
	fmt.Printf("of type %T\n", h)
}

func main() {
	fmt.Println("FUNCTIONS!")
	s := greeting("Hello", "Frank")
	fmt.Println(s)
	cm, degrees := complexNumber(1, 1)
	fmt.Println(cm, degrees)

	// variadic parameters - take an unlimited (zero or more) number of paramters of type T
	summation(1, 2, 3, 4, 5, 6, 7)
	// can also perform like this - spread the slice out:
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(summation(xi...))

	// defer - this will run the deferred function when the current function executes
	// eg: greeting will run last
	// use when you open a file, call defer file.Close() immediately afterwards
	fmt.Println("Defer Example: ")
	defer fmt.Println("\t", greeting("Hello", "Frank"))
	cm, degrees = complexNumber(0.5, 0.5*math.Sqrt(3))
	fmt.Printf("\t z = %v; with angle = %v\n", cm, degrees)

	// Methods: attach to structs.
	// interfaces/polymorphism: if a type has a method, then it is of that interface type as well
	// allows for functions to accept multiple types
	// fn bar can accpet any type that implements the speak method in this example (secretAgent and person)
	//     - func (r receiver) identifier(params) { // code }

	sa1 := secretAgent{
		person: person{
			first: "Branden",
			last:  "Smith",
		},
		ltk: true,
	}

	p1 := person{
		first: "James",
		last:  "Bond",
	}

	sa1.speak("I am alive")
	bar(sa1)
	bar(p1)
	printer(p1)

	// ananymous functions

	func(num int) {
		fmt.Println("anonymous function ran", num)
	}(15)

	// function expressions: first class citizens
	findMax := func(nums ...float64) float64 {
		max := math.Inf(-1)
		for _, num := range nums {
			if num > max {
				max = num
			}
		}
		return max
	}
	nums := []float64{1, 2, 3, 100, -50}
	fmt.Println("max number", findMax(nums...))

	calculator := map[string]func(x, y float64) float64{
		"+": func(x, y float64) float64 {
			return x + y
		},
		"-": func(x, y float64) float64 {
			return x - y
		},
		"*": func(x, y float64) float64 {
			return x * y
		},
		"/": func(x, y float64) float64 {
			return x / y
		},
	}
	sum := calculator["+"](1, 2)
	difference := calculator["-"](1, 2)
	product := calculator["*"](1, 2)
	divided := calculator["/"](1, 2)
	fmt.Println(sum, difference, product, divided)

	// returning a function
	adder := func(x int) func(int) int {
		z := "currying"
		return func(y int) int {
			fmt.Println(z)
			return x + y
		}
	}

	add5 := adder(5)
	fmt.Println(add5(10))

	// callbacks - functional programming - go does not promote it, but supports it.

	evenSum := evenSum(summation, []int{1, 2, 3, 4, 8, 12}...)
	fmt.Println(evenSum)

	// closure: create an incrementor
	a := incrementor()
	b := incrementor()
	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3
	fmt.Println(a()) // 4
	fmt.Println(b()) // 1
	fmt.Println(b()) // 2

	// recursion
	// factorial example
	fmt.Println(factorial(4))
}

/****************************
*
* functions used in main
*
****************************/

func factorial(n int) int {
	if n <= 1 {
		return n
	}
	return n * factorial(n-1)
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func greeting(greeting, st string) string {
	return fmt.Sprint(greeting, ", ", st)
}

// returning multiple values
func complexNumber(x, y float64) (string, float64) {
	return fmt.Sprintf("%v + %vi", x, y), math.Atan(y/x) * 180 / math.Pi
}

// variadic parameters - will store passed in args as a slice of type T - int in this case - similar to rest operator
func summation(x ...int) int {
	// fmt.Println("Variadic parameters: ")
	// fmt.Println("\t", x)
	// fmt.Printf("\t %T\n", x)

	sum := 0

	for _, val := range x {
		sum += val
	}

	return sum
}

func evenSum(f func(evenNums ...int) int, nums ...int) int {
	xi := make([]int, len(nums))

	for _, num := range nums {
		if num%2 == 0 {
			xi = append(xi, num)
		}
	}

	return f(xi...)
}
