package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

// type person struct {
// 	name string
// }
func main() {
	writingToFile()
	readingFromFile()
	errorsWithInfo()

	/*************************************
	***** regex and rune sample code *****
	**************************************/
	// x := "test this out 123"
	// x = "-123- test this out"
	// y := 'y'
	// fmt.Printf("Y is::: %T %v\n", y, y)
	// regex := regexp.MustCompile(`^-?\d+`)
	// fmt.Println(strconv.ParseInt(regex.FindString(x), 10, 0))

	// isMatch, err := regexp.MatchString(`-|\d`, "a")
	// fmt.Println(isMatch, err)
	// // fmt.Println(x)
	// for _, c := range x {
	// 	if c == '-' {
	// 		fmt.Println("yay")
	// 	} else if c == 'e' {
	// 		fmt.Println("EEE")
	// 	}
	// }
}

func writingToFile() {
	f, err := os.Create("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	content := strings.NewReader("Branden")
	_, err = io.Copy(f, content)
	if err != nil {
		log.Fatal(err)
	}
}

func readingFromFile() {
	f, err := os.Open("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}

func errorsWithInfo() {
	val, err := sqrt(-10)
	if err != nil {
		log.Printf("%T\n", err)
		log.Fatal(err)
	}
	log.Println(val)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Cannot take sqrt of negative number")
	}
	return math.Sqrt(x), nil
}
