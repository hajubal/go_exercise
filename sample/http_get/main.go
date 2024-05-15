package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//get call
	resp, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
