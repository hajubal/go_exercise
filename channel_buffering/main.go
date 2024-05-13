package main

import "fmt"

func main() {
	//2개의 메시지를 받을 수 있는 channel 생성, 두번째 파라미터가 버퍼의 갯수
	messages := make(chan string, 2)

	//channel 에 메시지를 보내므로 go 루틴 없이도 메시지를 보낼 수 있다.
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	channelRange()
}

// channel을 사용하지 않으면 go 루틴으로 송신 로직 후, 수신로직이 없으면 Lock이 발생한다.
func error_case() {
	c := make(chan int)
	c <- 1           //수신루틴이 없으므로 데드락
	fmt.Println(<-c) //코멘트해도 데드락 (별도의 Go루틴없기 때문)
}

func channelRange() {
	ch := make(chan int, 2)

	//채널에 송신
	ch <- 1
	ch <- 2

	//채널을 닫는다
	close(ch)

	//방법 1
	// 채널이 닫힌 것을 감지할 때까지 계속 수신
	/*
		for {
			if i, success := <-ch; success {
				println(i)
			} else {
				break
			}
		}
	*/

	// 방법 2
	// 위 표현과 동일한 채널 range 문
	for i := range ch {
		println(i)
	}
}
