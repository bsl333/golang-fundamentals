package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Ex 1: Launch two new go routines and use wait group for them to finish and return results")
	example1()

	fmt.Println("Ex 2: method sets")
	fmt.Print("\t")
	example2()

	fmt.Println("Ex 3: create a race condition via an incrementer program")
	example3()
}

func example1() {
	/******
	* Ex 1: in addition to the main goroutine, launch two additional goroutines
	* use waitgroups to ensure everything finishes before program exits
	******/

	var wg sync.WaitGroup
	const goRoutines int = 2

	wg.Add(goRoutines) // for two goroutines
	for i := 0; i < goRoutines; i++ {
		go func() {
			fmt.Println("\tGo routine:", runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
}

type person struct {
	name string
}

func (p *person) speak(message string) {
	fmt.Println("I,", p.name, ", said", message)
}

type human interface {
	speak(message string)
}

func saySomething(h human) {
	h.speak("This shall work")
}

func example2() {
	/******
	* Ex 2: Method sets
	* create a type person struct
	* attach a method speak to type person using a pointer receiver
	* create a type human interface
	* create func “saySomething”
	* show the following in your code
	*    - you CAN pass a value of type *person into saySomething
	*    - you CANNOT pass a value of type person into saySomething
	******/
	p := person{"Branden"}
	// saySomething(p) // won't work, as saySomething expects the pointer to p.
	saySomething(&p)

}

func example3() {
	/******
	* Ex 3: Using goroutines, create an incrementer program
	* read the incrementer value, store it in a new variable, yield the processor with runtime.Gosched(), increment the new variable
	* show that a race condition is raised
	******/

	counter := 0
	goRoutines := 50
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for ; goRoutines != 0; goRoutines-- {
		go func() {
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
			fmt.Println("counter", counter)
			wg.Done()
		}()
	}
	fmt.Println("Done")
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}
