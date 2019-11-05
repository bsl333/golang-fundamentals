package main

import "fmt"

func main() {
	// Arrays - static in size and size is part of its type; ie type = [len(arr)]dataType
	// e.g, type of var x [8]int = [8]int
	// can ONLY store values of the SAME TYPE

	var x [8]int
	fmt.Println(x)
	x[4] = 100
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Printf("%T\n", x)

	// slices - can ONLY store values of the SAME TYPE
	// slice sits on top of an array
	mySlice := []int{4, 5, 10, 23}
	fmt.Printf("%v\t%T\n", mySlice, mySlice)

	for i, val := range mySlice {
		fmt.Println(i, val)
	}

	// slice a slice - similar to python
	fmt.Println(mySlice[:])   // copy of mySlice
	fmt.Println(mySlice[1:3]) // will return index 1 and 2. should be written as mySlice[start: end)
	fmt.Println(mySlice[1:])  // grab element from 1 to end

	// append to slice - can append individual elements or another slice
	mySlice = append(mySlice, 1000)
	fmt.Println(mySlice)
	anotherSlice := []int{-90, 50}
	mySlice = append(mySlice, anotherSlice...)
	fmt.Println(mySlice)

	// deleting from a slice - use slice and append
	mySlice = append(mySlice[:3], mySlice[5:]...) // merge idxs from 0-2 and 5 til end - removed idx 3 & 4
	fmt.Println(mySlice)

	// using make to create slices - can set slice length and capacity of array it sits on top of
	// capacity is the max number of elements the slice can hold before the compiler has to reallocate more memory
	// for a new underlying array - ie more work the computer has to do, ineffecient.
	newSlice := make([]int, 10, 11) // makes a slice of length 10, filled with 0s, with a capacity of 11
	fmt.Println(len(newSlice), cap(newSlice))
	newSlice = append(newSlice, 100, 15) // appending two elements will cause the slice to exceed its capacity, so cap is doubled to 22
	fmt.Println(len(newSlice), cap(newSlice))

	// multidimensional slices - slices don't have to be of the same length
	slice1 := []string{"Branden", "Denise"}
	slice2 := []string{"Margot"}

	multiDimSlice := [][]string{slice1, slice2}
	fmt.Println(multiDimSlice)

	// Maps - key value stores (hash table)
	// if you enter a key that is not in the map, value returned is 0
	// the fix: val, ok := myMap[nonExistentKey] ok = false
	myMap := map[string]int{
		"Branden":   6617330225,
		"Denise":    6613131830,
		"Christine": 6614767262,
	}

	fmt.Println(len(myMap))

	if val, ok := myMap["Branden"]; ok { // val = 0, ok = true | false
		fmt.Println(val, "exists")
	}

	for key, val := range myMap {
		fmt.Printf("%v: %v\n", key, val)
	}

	// how to create a charmap in go - good for checking if something is anagram
	charMap := map[string]int{}
	for _, uniCode := range "These are the names that were provided" {
		char := string(uniCode)
		if char == " " {
			continue
		}
		charMap[char]++
	}
	fmt.Println(charMap)

	// deleting from a map - use Delete function (built in)

	delete(myMap, "Denise")
	fmt.Println(myMap)
}
