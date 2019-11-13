package main

import "fmt"

func main() {
	exercise1()
	exercise2()
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
