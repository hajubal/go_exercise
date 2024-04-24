package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://www.google.com/robots.txt")

	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)

	fmt.Println("Response Status:", response.Status)

	scanner := bufio.NewScanner(response.Body)
	//scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

}
