package main

import (
	"course-material/section11-application/getEx"
	"encoding/json"
	"fmt"
)

// encode to JSON struct
type encodePersonToJson struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

//decode to jsonStruct
type decodeJsonToPerson struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Age       int    `json:"age"`
}

func main() {
	getEx.RetrieveUsers()
	fmt.Println("Marshel and Unmarshel")

	p1 := encodePersonToJson{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := encodePersonToJson{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []encodePersonToJson{p1, p2}
	fmt.Println(people)
	/******
	*  Marsheling: Convert struct to JSON
	******/
	jsonBytes, err := json.Marshal(people) // will return a slice of bytes ==> convert to string to get readable output
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonBytes))

	/******
	*  Marsheling: Convert struct to JSON
	******/
	// slice of bytes of uint8
	byteSlice := []byte(`[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`)

	var persons []decodeJsonToPerson
	err = json.Unmarshal(byteSlice, &persons)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(persons)

	for _, person := range persons {
		fmt.Println(person.FirstName, person.LastName, person.Age)
	}

}
