package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d staring\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			// defer 는 함수가 리턴되기 직전에 호출됨, 함수내에서 패닉이 발생되어도 실행, 시작부분에 있어서 가독성 향상
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
