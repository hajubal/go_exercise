package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with int"}
	}
	return arg + 3, nil
}

func main() {

	_, err := f(42)
	var ae *argError

	fmt.Println("ae >> ", ae)
	fmt.Println("&ae >> ", &ae)

	var ae2 argError

	fmt.Println("ae2 >> ", ae2)
	fmt.Println("&ae2 >> ", &ae2)

	ae3 := argError{1, "message"}

	fmt.Println("ae3 >> ", ae3)
	fmt.Println("&ae3 >> ", &ae3)

	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}

}
