package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First string `json:"first"`
	Age   int    `json:"age"`
}

type detailedUser struct {
	user
	Last    string
	Sayings []string
}

func (du detailedUser) String() string {
	return fmt.Sprintf(`
	First: %s
	Last: %s
	Age: %d
	Sayings: %v
	`, du.First, du.Last, du.Age, du.Sayings)
}

type byAge []detailedUser

func (ba byAge) Len() int           { return len(ba) }
func (ba byAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }
func (ba byAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }

type byLast []detailedUser

func (bl byLast) Len() int           { return len(bl) }
func (bl byLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }
func (bl byLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }

func main() {

	/******
	* Ex 1: Marshal the structs to convert to JSON
	******/
	fmt.Println("Ex 1: Marshaling")

	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println("\tPre marshaling\n\t\t", users)

	// your code goes here
	bytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\tPost marshaling\n\t\t", string(bytes))

	fbytes, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println(string(fbytes))

	/******
	* Ex 2: Unmarshal JSON to a struct
	******/
	fmt.Println("Ex 2: Unmarshaling")

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	// fmt.Println(s)
	var detailedUsers []detailedUser
	if err := json.Unmarshal([]byte(s), &detailedUsers); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\t", detailedUsers)

	/******
	* Ex 3: encode a struct to JSON and write to stdout
	* used detailedUses from last example
	******/
	fmt.Println("Ex 3: Using json encode to wrote to stdout")

	if err := json.NewEncoder(os.Stdout).Encode(detailedUsers); err != nil {
		fmt.Println(err)
		return
	}

	/******
	* Ex 4: sort the []int and []string
	******/
	fmt.Println("Ex 4: Basic sorting")

	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println("\t", xi)
	sort.Ints(xi)
	fmt.Println("\t", xi)

	fmt.Println("\t", xs)
	sort.Strings(xs)
	fmt.Println("\t", xs)

	/******
	* Ex 5: sort detailedUser by Age and Last
	* also sort the Sayings of each user
	******/
	fmt.Println("Ex 5: Custom sorting")

	// sorting sayings first
	for _, du := range detailedUsers {
		sort.Strings(du.Sayings)
	}
	sort.Sort(byAge(detailedUsers))
	fmt.Println("\tbyAge\n", detailedUsers)
	sort.Sort(byLast(detailedUsers))
	fmt.Println("\tbyAge\n", detailedUsers)
}
