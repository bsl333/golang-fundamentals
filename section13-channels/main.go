package main

import (
	"fmt"
	"runtime"
)

func main() {
	channelBasics()
	directionalChannels()
	channelApp()
	rangeOverChannel()
	usingSelect()

}

// bidirectional channel (basics)
// can send to and receive from a channel
func channelBasics() {
	c := make(chan int)
	go func() { // fire off another go routine
		c <- 42 // put 42 on c
	}()

	fmt.Println(<-c) // read value from channel.

	// or can make a buffered channel where you allocate space for whatever you're putting on
	c = make(chan int, 1)

	c <- 42
	fmt.Println(<-c)

}

// directional channels
func directionalChannels() {
	// can only put things on the channel (send only)
	sendOnlyCh := make(chan<- int)
	fmt.Printf("%T\n", sendOnlyCh)
	receiveOnlyCh := make(<-chan int)
	fmt.Printf("%T\n", receiveOnlyCh)
}

func sendToCh(ch chan<- int, val int) {
	ch <- val
}

func receiveFromCh(ch <-chan int) {
	fmt.Println(<-ch)
}

func channelApp() {
	ch := make(chan int)
	for i := 0; i < 100; i++ {
		go sendToCh(ch, i)
		receiveFromCh(ch)
	}
}

func rangeOverChannel() {
	ch := make(chan int)

	go func(c chan<- int) {
		for i := 0; i < 10; i++ {
			sendToCh(ch, i)
			println("sending", i)
		}
		close(ch)
	}(ch)

	fmt.Println("go routines", runtime.NumGoroutine())

	for val := range ch {
		fmt.Println("pulling", val)
	}
}

// can fire off go routines and write to channels
// then use select statement to perform some action when those channels are ready to be writen from.
func usingSelect() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan string)

	// send to channel
	go func(even, odd chan<- int, quit chan<- string) {
		for i := 0; i < 20; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		// close(quit)
		quit <- "Done"
		close(odd)
		close(even)
	}(even, odd, quit)

	for {
		select {
		case num := <-even:
			fmt.Println("Even number:", num)
		case num := <-odd:
			fmt.Println("Odd num:", num)
		case msg := <-quit:
			fmt.Println(msg)
			return
		}
	}
}
