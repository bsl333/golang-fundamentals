package main

import (
	"fmt"
	"math"
)

type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Printf("\tI am %v %v and am %v years old\n", p.first, p.last, p.age)
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

type shape interface {
	area() float64
}

func info(shape shape) {
	switch shape.(type) {
	case circle:
		fmt.Println("\t circle area = ", shape.area())
	case square:
		fmt.Println("\t square area = ", shape.area())
	default:
		fmt.Println("\t unknow shape")
	}
}

type vector []interface{}

func (v vector) forEach(cb func(val interface{}, i int)) vector {
	for i, val := range v {
		cb(val, i)
	}
	return v
}

func main() {
	/******
	 * Ex 1: create the following simple functions
	 * create a func with the identifier foo that returns an int
	 * create a func with the identifier bar that returns an int and a string
	* call both funcs
	* print out their results
	******/
	fmt.Println("Ex 1: Simple funtions")
	foo := func() int {
		return 3
	}
	bar := func() (int, string) {
		return 3, "three"
	}
	fmt.Println("\t", foo())
	num, str := bar()
	fmt.Println("\t", num, str)

	/******
	 * Ex 2: variadic parameters vs slice parameters
	 * create a func with the identifier sumVariadic that
	 * takes in a variadic parameter of type int
	 * pass in a value of type []int into your func (unfurl the []int)
	 * returns the sum of all values of type int passed in
	 * create a func with the identifier sumSlice that
	 * takes in a parameter of type []int
	 * returns the sum of all values of type int passed in
	 ******/
	fmt.Println("Ex 2: Using variadic params and slice params")
	nums := []int{1, 2, 3, 4, 5}
	sumVariadic := func(nums ...int) int {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		return sum
	}
	sumSlice := func(nums []int) int {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		return sum
	}
	fmt.Println("\t", sumVariadic(nums...))
	fmt.Println("\t", sumSlice(nums))
	/******
	* Ex 3: Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
	******/

	defer fmt.Println("Ex 3: Use the defer keyword\n\t I am deferred!")

	/******
	* Ex 4: Create a user defined struct with the identifier “person” the fields:
	* first, last, age
	* attach a method to type person with the identifier “speak”
	* the method should have the person say their name and age
	* create a value of type person and call the method from the value of type person
	******/
	fmt.Println("Ex 4: Attach a method to a struct")
	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p.speak()

	/******
	* Ex 5: create a type SQUARE; create a type CIRCLE
	* attach a method to each that calculates AREA and returns it
	* circle area= π r^2
	* square area = L * W
	* create a type SHAPE that defines an interface as anything that has the AREA method
	* create a func INFO which takes type shape and then prints the area
	* create a value of type square; create a value of type circle
	* use func info to print the area of square and the square
	******/
	fmt.Println("Ex 5: Interfaces ")
	c := circle{
		radius: 3,
	}
	sq := square{
		length: 5,
	}

	info(c)
	info(sq)

	/******
	* Ex 6: Build and use an anonymous func
	******/
	fmt.Println("Ex 6: Use an anonymous function")
	func(punctuation string) {
		fmt.Println("\t I am an anonymous function" + punctuation)
	}("!!!")

	/******
	* Ex 7: assign a function to a variable
	******/
	fmt.Println("Ex 7: Assign a function to a variable")
	greeting := func(message string) {
		fmt.Println("\t", message)
	}
	greeting("Hi there!")

	/******
	* Ex 8: Create a func which returns a func
	******/
	fmt.Println("Ex 8: Functions that return functions")

	greeter := func(message string) func(string) string {
		return func(name string) string {
			return fmt.Sprintf("%v, %v?", message, name)
		}
	}

	meanGreeting := greeter("What do you want")
	fmt.Println(meanGreeting("James"))
	niceGreeting := greeter("Hi! How are you")
	fmt.Println(niceGreeting("Denise"))

	/******
	* Ex 9: A “callback” is when we pass a func into a func as an argument.
	******/
	fmt.Println("Ex 9: Use a callback - implementing forEach")
	v := vector{1, 4, 6, 10, -20, 30, 4, 4, 10}
	v2 := vector{"john", "bob", "john", "dan"}
	mapCounter := map[interface{}]int{}
	cb := func(val interface{}, idx int) {
		mapCounter[val]++
	}
	v.forEach(func(val interface{}, idx int) {
		mapCounter[val]++
	})
	v2.forEach(cb)
	fmt.Println("\t", mapCounter)

	/******
	* Ex 10: Closure is when we have “enclosed” the scope of a variable in some code block.
	******/
	fmt.Println("Ex 10: Use closure: build an incrementer")

	incrementer := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}

	counter := incrementer()
	fmt.Println("\t", counter())
	fmt.Println("\t", counter())
	fmt.Println("\t", counter())
	fmt.Println("\t", counter())

}
