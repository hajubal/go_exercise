package main

import (
	"encoding/json"
	"fmt"
)

// Member
type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {

	/**
	인코딩
	*/
	mem := Member{"Alex", 10, true}

	//json encoding
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	//json byte를 문자열로 변경
	fmt.Println(string(jsonBytes))

	/**
	디코딩
	*/
	jsonBytes, _ = json.Marshal(Member{"Time", 40, false})
	err = json.Unmarshal(jsonBytes, &mem)
	if err != nil {
		panic(err)
	}

	fmt.Println(mem)

}
