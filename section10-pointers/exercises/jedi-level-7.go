package main

import "fmt"

type person struct {
	name    string
	address string
}

func changeMe(person *person, newAddress string) {
	(*person).address = newAddress
}

func main() {
	/******
	* Ex 1: create a variable and assign value; print its address
	******/
	fmt.Println("Ex 1: Print a variables address")
	x := 3
	fmt.Println("\t", &x)

	/******
	* Ex 2: create a person struct; create a func called “changeMe” with a *person as a parameter that changes the address
	******/
	fmt.Println("Ex 2: create a function to change a person's address field, must use pointers")
	p := person{
		name:    "Donald Trump",
		address: "1600 Pennsylvania Ave SE Washington, DC 20003",
	}

	fmt.Printf("\t%v was at %v\n", p.name, p.address)
	changeMe(&p, "1901 D St E Washington, DC 20003")
	fmt.Printf("\t%v now at %v\n", p.name, p.address)
}
