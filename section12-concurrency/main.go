package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// sampleWaitGroup()
	mockRaceCondition()
}

var wg sync.WaitGroup

func sampleWaitGroup() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("goroutines\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func mockRaceCondition() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)
	counter := 0

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			// v := counter
			a := &counter
			// time.Sleep(time.Millisecond * 10)
			// runtime.Gosched()
			// v++
			// counter = v
			// counter++
			*a++
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("Count:", counter)

}
