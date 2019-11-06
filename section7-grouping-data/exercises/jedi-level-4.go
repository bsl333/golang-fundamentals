package main

import "fmt"

func main() {
	/*****
	* Exercise 1: Create an array of 5 ints, print values and array type
	*****/
	fmt.Println("Ex 1")

	myArr := [5]int{2, 3, 5, 7, 11}
	for _, v := range myArr {
		fmt.Println("\t", v)
	}
	fmt.Printf("\t %T\n", myArr)

	/*****
	* Exercise 2: Create an slice of 10 ints, print values and slice type
	*****/

	fmt.Println("Ex 2")
	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, v := range mySlice {
		fmt.Println("\t", v)
	}
	fmt.Printf("\t %T\n", mySlice)

	/*****
	* Exercise 3: Using previous slice, slice to creat the following slices:
	*	[42 43 44 45 46]
	*	[47 48 49 50 51]
	*	[44 45 46 47 48]
	*	[43 44 45 46 47]
	*****/

	fmt.Println("Ex 3")
	fmt.Println("\t", mySlice[:5])
	fmt.Println("\t", mySlice[5:])
	fmt.Println("\t", mySlice[2:7])
	fmt.Println("\t", mySlice[1:6])

	/*****
	* Exercise 4: append to mySlice:
	* 52
	* 53, 54, 55
	* y := []int{56, 57, 58, 59, 60}
	*****/

	fmt.Println("Ex 4")

	mySlice = append(mySlice, 52)
	fmt.Println("\t", mySlice)

	mySlice = append(mySlice, 53, 54, 55)
	fmt.Println("\t", mySlice)

	// otherSlice := []int{56, 57, 58, 59, 60}
	// mySlice = append(mySlice, otherSlice...)
	mySlice = append(mySlice, []int{56, 57, 58, 59, 60}...)
	fmt.Println("\t", mySlice)

	/*****
	* Exercise 5: Delete from the slice x, and create a new slice, y, which has the following values:
	* y = [42, 43, 44, 48, 49, 50, 51]
	*****/
	fmt.Println("Ex 5")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println("\t", y)

	/*****
	* Exercise 6: Create a slice to store the names of all of the states in the United States of America
	* use the make func
	* print all states and idx without using range clause
	*****/
	fmt.Println("Ex 6")
	states := make([]string, 0, 50)

	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	for i := 0; i < len(states); i++ {
		fmt.Println("\t", i, states[i])
	}

	/*****
	* Exercise 7: Create a multidimensional slice with the following slices
	* ["James", "Bond", "Shaken, not stirred"]
	* ["Miss", "Moneypenny", "Helloooooo, James."]
	*****/
	fmt.Println("Ex 7")

	mutliDimSlice := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, xi := range mutliDimSlice {
		fmt.Println("\t", xi)
		for _, v := range xi {
			fmt.Println("\t\t", v)
		}
	}

	/*****
	* Exercise 8: Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string
	* which stores their favorite things. Store three records in your map.
	*****/
	fmt.Println("Ex 8")
	favoriteThingsMap := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for name, favoriteThings := range favoriteThingsMap {
		fmt.Println("\tFor record", name)
		for _, favoriteThing := range favoriteThings {
			fmt.Println("\t\t", favoriteThing)
		}
	}

	/*****
	* Exercise 9: Add something new to the map, then print using range loop
	*****/
	fmt.Println("Ex 9")
	favoriteThingsMap["lowe_branden"] = []string{"Basketball", "Coding", "Eating"}
	for name, favoriteThings := range favoriteThingsMap {
		fmt.Printf("\t%v\n\t\t%v\n", name, favoriteThings)
	}

	/*****
	* Exercise 10: Delete something from the map, then print using range loop
	*****/
	fmt.Println("Ex 10")

	delete(favoriteThingsMap, "no_dr")
	for name, favoriteThings := range favoriteThingsMap {
		fmt.Printf("\t%v\n\t\t%v\n", name, favoriteThings)
	}

}
