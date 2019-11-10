package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// need to rewatch lecture, but essentially, it is showing how interfaces make the standard library
	// much more usable/powerful
	// type File has a method Write (b []byte)(n int, err error)
	//	- this is the same function signature for a write interface
	//  - so any method that accepts a Writer also accepts a *File

	// function signature:  io.WriteString(w Writer, s string) (n int, err error)
	// Stdout creates a new File of type *File, which is also of type Write, hence why the below code functions
	if _, err := io.WriteString(os.Stderr, "Hello, Playground\n"); err != nil {
		fmt.Println(err)
	}

	// example for user input

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		fmt.Println("user input:", input)
	}

	if scanner.Err() != nil {
		fmt.Println("error")
	}

}
