package main

import "fmt"

type person struct {
	first string // IDENTIFIER and TYPE
	last  string
	age   int
}

// embedded struct
type secretAgent struct {
	person      // anonymous field name
	ltk    bool // license to kill
}

func main() {

	// p1 is a value of type person - NOT an object or class
	p1 := person{
		first: "Branden",
		last:  "Lowe",
		age:   28,
	}

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)

	// embedded structs -- inner structs FIELDS get PROMOTED to OUTER TYPE, so you can access them directly with dot notation
	// inner struct = person; outer struct = secret agent.
	// if there are name collisions, where inner struct has same field name as outer struct, can access by:
	// outerStruct.innerStruct.prop
	sa1 := secretAgent{
		person: person{
			first: "Sam",
			last:  "Smith",
			age:   100,
		},
		ltk: true,
	}
	fmt.Println(sa1.first, sa1.last, sa1.ltk)

	// anonymous struct
	house := struct {
		sqft         int
		numBedrooms  int
		numBathrooms float64
		state        string
	}{
		sqft:         1200,
		numBedrooms:  3,
		numBathrooms: 1.5,
		state:        "CA",
	}
	fmt.Println(house.sqft, house.state)
}
