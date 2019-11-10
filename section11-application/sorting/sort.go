package main

import (
	"fmt"
	"sort"
	"strings"
)

type person struct {
	First string
	Age   int
}

type byAge []person
type byName []person

func (ba byAge) Len() int           { return len(ba) }
func (ba byAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }
func (ba byAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }

func (bn byName) Len() int { return len(bn) }
func (bn byName) Less(i, j int) bool {
	return strings.ToUpper(bn[i].First) < strings.ToUpper(bn[j].First)
}
func (bn byName) Swap(i, j int) { bn[i], bn[j] = bn[j], bn[i] }

func main() {
	basicSorting()
	customSorting()
}

func (p person) String() string {
	return fmt.Sprintf(`
	"name": %s
	"age": %d
	`, p.First, p.Age)
}

func customSorting() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}
	p5 := person{"N", 94}
	p6 := person{"m", 1}

	fmt.Println("A" > "a")

	people := []person{p1, p2, p3, p4, p5, p6}
	fmt.Println(people)

	sort.Sort(byAge(people))
	fmt.Println("Sorted by age:\n", people)
	sort.Sort(byName(people))
	fmt.Println("Sorted by name:\n", people)
}

func basicSorting() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	sort.Ints(xi)
	fmt.Println(xi)
	sort.Sort(sort.Reverse(sort.IntSlice(xi)))
	fmt.Println(xi)
	for _, str := range xs {
		for i := 0; i <= len(str); i++ {
			fmt.Print("-")
		}
	}
	fmt.Println("-")
	sort.Strings(xs)
	fmt.Println(xs)
}
