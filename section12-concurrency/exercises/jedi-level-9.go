package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Ex 1: Launch two new go routines and use wait group for them to finish and return results")
	example1()
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
