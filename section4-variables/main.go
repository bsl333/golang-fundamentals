package main

import (
	"fmt"
)

type baller float64

var vc baller

func main() {
	x := 20.0
	fmt.Println("Hello world")

	vc = 15.23

	fmt.Printf("%v\n", vc)
	fmt.Printf("%T\n", vc)
	x = float64(vc)
	// resp, err := http.Get("http://jsonplaceholder.typicode.com/posts")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// stringified := string(body)
	// fmt.Println(stringified)
	fmt.Println(x)
}
