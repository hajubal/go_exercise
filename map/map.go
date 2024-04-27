package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("delete map:", m)

	clear(m)
	fmt.Println("clear map:", m)

	// 맵(map) m에서 키 "k2"에 해당하는 값을 조회하고, 그 결과를 두 개의 변수에 할당하는 방식입니다.
	// 이 구문에서 첫 번째 반환 값은 맵의 해당 키에 저장된 값입니다. 하지만 여기서는 첫 번째 반환 값을 무시하고 _ (밑줄)을 사용해 그 값을 무시합니다.
	// 이는 해당 값이 필요하지 않을 때 자주 사용되는 방식입니다.
	//
	// 두 번째 반환 값인 prs는 해당 키가 맵에 존재하는지 여부를 나타내는 불리언(Boolean) 값입니다.
	// 키 "k2"가 맵에 존재하면 prs는 true가 되고, 그렇지 않다면 false가 됩니다.
	// 이 값은 키의 존재 여부만 확인할 때 유용하게 사용됩니다.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
