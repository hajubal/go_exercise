package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {

	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 22},
		Person{name: "Smith", age: 32},
		Person{name: "Jain", age: 52},
	}

	slices.SortFunc(people, func(i, j Person) int {
		return cmp.Compare(i.age, j.age)
	})

	fmt.Println(people)
}
