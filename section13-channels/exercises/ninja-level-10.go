package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	/******
	* Fix the below code
	******/
	fmt.Println("Ex 1: fix channel implementation so code runs.")
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println("\t", <-c)
}

func exercise2() {
	/******
	* Fix the below code
	******/
	fmt.Println("Ex 2: fix channel implementation so code runs.")
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println("\t", <-cs)
	fmt.Printf("\t cs%T\n", cs)
}

func exercise3() {
	/******
	* Ex 3: use the range clause to print off values from channel
	******/
	fmt.Println("Ex 3: Range clause to pull values off of channel")
	c := gen()
	receive(c)

	fmt.Println("\tabout to exit")
}
func receive(ch <-chan int) {
	for val := range ch {
		fmt.Println("\t", val)
	}
}

func gen() <-chan int {
	c := make(chan int)
	go func(c chan<- int) {
		for i := 0; i < 15; i++ {
			c <- i
		}
		close(c)
	}(c)

	return c
}
