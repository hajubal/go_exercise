package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println("\nsp := &s")
	fmt.Printf("sp.age: %d\n", sp.age)
	fmt.Printf("s.age: %d\n", s.age)

	sp.age = 51
	fmt.Println("\nsp.age = 51")
	fmt.Printf("sp.age: %d\n", sp.age)
	fmt.Printf("s.age: %d\n", s.age)

	sp2 := s
	fmt.Println("\nsp2 := s")
	fmt.Printf("sp2.age: %d\n", sp2.age)
	fmt.Printf("s.age: %d\n", s.age)

	sp2.age = 53
	fmt.Println("\nsp2.age = 53")
	fmt.Printf("sp2.age: %d\n", sp2.age)
	fmt.Printf("s.age: %d\n", s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
