package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
}

func example1() {
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

func example2() {
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

func example3() {
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

func example4() {
	/******
	* Ex 4: use a select statement to pull values off of channel
	******/
	fmt.Println("Ex 4: Read channel values using select statement")

	q := make(chan int)
	c := gen2(q)

	receive2(c, q)
}

func receive2(c, q <-chan int) {
	for {
		select {
		case num := <-c:
			fmt.Println("\tvalue from c:", num)
		case <-q:
			fmt.Println("\tClosing")
			return
		}
	}
}

func gen2(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 15; i++ {
			c <- i
		}
		q <- 1
		close(c)
		close(q)
	}()

	return c
}

func example5() {
	/******
	* EX 5: use the val, ok idiom to print value from channel
	******/
	fmt.Println("Ex 5: ok idiom when pulling values off of a channel")
	c := make(chan int)
	go func() {
		c <- 3
		close(c)
	}()
	if v, ok := <-c; ok { // will print
		fmt.Println("\t", v, ok)
	}
	if v, ok := <-c; ok { // won't print
		fmt.Println(v, ok)
	}
}

func example6() {
	/******
	* Ex 6: Write a program that puts 100 numbers to a channel and pull the numbers off the channel and print them
	******/
	fmt.Println("Ex 6: Write a program to write 100 nums to channel, then read them from channel")

	ch := make(chan int)

	go func(ch chan<- int) {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	for val := range ch {
		fmt.Println("\t", val)
	}
}

func example7() {
	/******
	* Ex 7: write a program that
	* 	- launches 10 goroutines
	* 	- each goroutine adds 10 numbers to a channel
	* 	- pull the numbers off the channel and print them
	******/
	fmt.Println("Ex 5: launch 10 go routines to read to a channel, then read from that channel")
	ch := make(chan int)
	const goRoutines int = 5

	// create 10 go routines
	for gr := 0; gr < goRoutines; gr++ {
		// each go routine adds 10 unique numbers
		go func(gr int, ch chan<- int) {
			i := gr * 10
			for ; i < (gr+1)*10; i++ {
				ch <- i
			}
		}(gr, ch)
	}

	for i := 0; i < 10*goRoutines; i++ {
		fmt.Println("\t", i, "\t", <-ch)
	}

}
