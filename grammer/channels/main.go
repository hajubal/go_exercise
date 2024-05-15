package main

import "fmt"

// 채널은 흔히 goroutine들 사이 데이타를 주고 받는데 사용되는데,
// 상대편이 준비될 때까지 채널에서 대기함으로써 별도의 lock을 걸지 않고 데이타를 동기화하는데 사용
func main() {
	//channel 생성
	messages := make(chan string)

	//goroutine을 이용해서 messages channel에 데이터 전송
	go func() { messages <- "ping" }()

	//수신된 데이터 주입
	msg := <-messages

	fmt.Println(msg)
}
