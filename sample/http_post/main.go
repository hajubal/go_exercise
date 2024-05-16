package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	//간단한 http.post 예제
	reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Response 체크
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(respBody))
	}
}
