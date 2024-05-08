package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine 11")
	go f("goroutine 22")
	go f("goroutine 33")
	go f("goroutine 44")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second * 3)
	fmt.Println("done")
}
