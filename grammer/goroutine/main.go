package main

import (
	"fmt"
	"runtime" //multi cpu
	"sync"
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

	waitTest()
}

func waitTest() {
	//WaitGroup 2개 생성, 2개의 Go 루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2)

	fmt.Println("first go routine")
	go func() {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println("goroutine 1")
	}()

	fmt.Println("second go routine")
	go func(msg string) {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println("goroutine 2", msg)
	}("Hi")

	fmt.Println("waiting")
	wait.Wait() //go 루틴이 끝날 때까지 대기
}

/*
*
Go는 디폴트로 1개의 CPU를 사용한다.
즉, 여러 개의 Go 루틴을 만들더라도, 1개의 CPU에서 작업을 시분할하여 처리한다 (Concurrent 처리).
만약 머신이 복수개의 CPU를 가진 경우, Go 프로그램을 다중 CPU에서 병렬처리 (Parallel 처리)하게 할 수 있는데,
병렬처리를 위해서는 아래와 같이 runtime.GOMAXPROCS(CPU수) 함수를 호출하여야 한다
(여기서 CPU 수는 Logical CPU 수를 가리킨다).
*/
func multiCpuTest() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//...
}
