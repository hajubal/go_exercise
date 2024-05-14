package main

import (
	"fmt"
	"time"
)

/*
*
- timer: 일정시간 이후에 한번만 실행
- ticker: 일정시간 동안 반복적으로 실행
*/
func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(5000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
