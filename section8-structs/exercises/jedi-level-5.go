package main

import "fmt"

func printSlice(person person) {
	fmt.Println("\t", person.first, "favorite ice creams:")
	for _, v := range person.favoriteIceCreams {
		fmt.Println("\t\t", v)
	}
}

type person struct {
	first, last       string
	favoriteIceCreams []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	fmt.Println()
	/*****
	* Exercise 1: Create your own type “person”
	* has fields; first, last, favorite ice cream flavors ([]string)
	* create two values of type person
	* print favorite icream flavors
	*****/

	fmt.Println("Ex 1")
	p1 := person{
		first: "Branden",
		last:  "Smith",
		favoriteIceCreams: []string{
			"Vanilla",
			"Salted Caramel",
			"Cookie Dough",
		},
	}

	p2 := person{
		first: "Margot",
		last:  "Jenkins",
		favoriteIceCreams: []string{
			"Chocolate",
			"Fudge",
			"Strawberry",
		},
	}

	printSlice(p1)
	printSlice(p2)

	/*****
	* Exercise 2:Take the code from the previous exercise, then store the values
	* of type person in a map with the key of last name. Access each value in the
	* map. Print out the values, ranging over the slice.
	*****/

	personMap := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range personMap {
		fmt.Println(v)
	}

	/*****
	* Exercise 3: create a new type vehicle with fields, doors, color
	* create two new types truck, with an additional field fourWheel drive; sedan with additional field luxury
	* using composite literals, create variables of type truck and sedan; print each out, then print one field from each out.
	*****/
	fmt.Println("Ex 3")

	truck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Black",
		},
		fourWheel: true,
	}

	sedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "White",
		},
		luxury: true,
	}

	fmt.Println("\tTruck:", truck, "has", truck.doors, "doors")
	fmt.Println("\tSedan:", sedan, "is", sedan.color)

	/*****
	* Exercise 4: Create and use an anonymous struct
	* will use one for a computer
	*****/

	fmt.Println("Ex 4")

	myComputer := struct {
		brand    string
		cpu      string
		cores    int
		ram      int
		isLaptop bool
	}{
		brand:    "Apple",
		cpu:      "Intel i5",
		cores:    2,
		ram:      16,
		isLaptop: true,
	}

	fmt.Println("\t my laptop specs are", myComputer)

}
