package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha 256 string test"
	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
