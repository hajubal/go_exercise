package main

import (
	"fmt"
	"slices"
)

func main() {

	strs := []string{"b", "d", "a"}
	slices.Sort(strs)
	fmt.Println("strs:", strs)

	ints := []int{10, 5, 3, 4, 2, 7, 6, 9, 8, 1}
	slices.Sort(ints)
	fmt.Println("ints:", ints)

	fmt.Println(slices.IsSorted(ints))
}
