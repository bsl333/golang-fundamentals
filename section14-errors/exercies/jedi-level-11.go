package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	exercise1()
	exercise2()
}

type customErr struct {
	code string
	msg  string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("code: %v: %s", ce.code, ce.msg)
}

func foo(err error) {
	fmt.Println(err)
}

func exercise1() {
	/******
	* Create a struct “customErr” which implements the builtin.error interface.
	* Create a func “foo” that has a value of type error as a parameter.
	* Create a value of type “customErr” and pass it into “foo”
	******/
	fmt.Println("Ex 1: create a struct that implements the builtin error interface")
	e := customErr{"404", "Not Found"}
	foo(e)
	fmt.Println()
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (sqErr sqrtError) Error() string {
	return fmt.Sprintf("lat: %s, long: %s. %v", sqErr.lat, sqErr.long, sqErr.err)
}

func exercise2() {
	/******
	* Use the sqrt.Error struct as a value of type error. If you would like, use these numbers for your
	******/
	fmt.Println("Ex 2:")
	val, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		errMsg := fmt.Errorf("Error in sqrt, cannot calculate square root of negative number. Invalid input: %v", num)
		return 0, sqrtError{"75.5", "87.2", errMsg}
	}

	return math.Sqrt(num), nil
}
