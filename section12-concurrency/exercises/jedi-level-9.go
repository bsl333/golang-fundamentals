package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Ex 1: Launch two new go routines and use wait group for them to finish and return results")
	example1()

	fmt.Println("Ex 2: Method sets")
	fmt.Print("\t")
	example2()

	fmt.Println("Ex 3: Create a race condition via an incrementer program")
	example3()

	fmt.Println("Ex 4: Fix example 3 to eliminate race conditions")
	example4()

	fmt.Println("Ex 5: Fix example 3 using atomic package")
	example5()

	fmt.Println("Ex 6: Create a program that prints out your OS and ARCH.")
	example6()
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
			// runtime.Gosched()
			temp++
			counter = temp
			fmt.Println("counter", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("Done")
	fmt.Println(runtime.NumGoroutine())
}

func example4() {
	/******
	* Ex 4: fix example 3 with mutexes to prevent race conditions
	******/
	counter := 0
	goRoutines := 50
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(goRoutines)

	for ; goRoutines != 0; goRoutines-- {
		go func() {
			mutex.Lock()
			temp := counter
			temp++
			counter = temp
			fmt.Println("counter", counter)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Done")
	fmt.Println(runtime.NumGoroutine())
}

func example5() {
	/******
	* Ex 5: fix ex 3 via atomic package
	******/
	var counter int64
	goRoutines := 10
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for ; goRoutines != 0; goRoutines-- {
		go func() {
			atomic.AddInt64(&counter, 1) // will increment counter by 1
			runtime.Gosched()
			fmt.Println("counter", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func example6() {
	/******
	* Ex 6: Create a program that prints out your OS and ARCH. Use the following commands to run it
	* go run, build, install
	******/

	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
}
