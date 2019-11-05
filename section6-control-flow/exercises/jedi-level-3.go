package main

import "fmt"

func main() {
	/****
	* Exercise 1: print every number from 1 thru 10,000
	****/
	fmt.Println("Ex 1")
	for i := 1; i <= 10000; i++ {
		fmt.Println("\t", i)
	}

	/****
	* Exercise 2: Print every rune code point of the uppercase alphabet three times.
	****/
	fmt.Println("Ex 2")
	for i := 65; i <= 90; i++ {
		fmt.Println("\t", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}

	/****
	* Exercise 3: Print out the years you have been alive; use for conidtion {} style loop
	****/
	fmt.Println("Ex 3")
	year := 1991
	for year <= 2019 {
		fmt.Println("\t", year)
		year++
	}

	/****
	* Exercise 4: Print out the years you have been alive; use for {} style loop
	****/
	fmt.Println("Ex 4")
	year = 1991
	for {
		if year > 2019 {
			break
		}
		fmt.Println("\t", year)
		year++
	}

	/****
	* Exercise 5: Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
	****/
	fmt.Println("Ex 5")
	for i := 10; i <= 100; i++ {
		fmt.Printf("\t the remainder of %v divided by 4 is %v\n", i, i%4)
	}

	/****
	* Exercise 6/7: if else statements
	****/

	num := 3
	fmt.Println("Ex 6 & 7")
	if num < 0 {
		fmt.Println("\t", "num is less than 0")
	} else if num < 5 {
		fmt.Println("\t", "num is greater than 0 but less than 5")
	} else {
		fmt.Println("\t", "num is greater than 5")
	}

	/****
	* Exercise 8: switch statement with no switch expression
	****/
	fmt.Println("Ex 8")
	name := "Branden"
	switch {
	case name == "Mike":
		fmt.Println("\t", "Name is Mike")
	case name == "Branden":
		fmt.Println("\t", "Name is Branden")
	default:
		fmt.Println("\t", "IDK your name bro")
	}

	/****
	* Exercise 9: switch statement with a switch expression of TYPE string with the IDENTIFIER “favSport”
	****/
	fmt.Println("Ex 9")
	favSport := "Basketball"
	switch favSport {
	case "Soccer":
		fmt.Println("\t", "Soccer is your fav sport")
	case "Basketball":
		fmt.Println("\t", "Basketball is your fav sport")
	default:
		fmt.Println("\t", "IDK your fav sport")
	}
}
