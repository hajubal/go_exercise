package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//worker pool 생성
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//worker pool에 작업 갯수 만큼 전달
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	//채널이 닫히면 더이상 데이터를 받을 수 없음
	close(jobs)

	//결과 받기, 이게 없으면 프로세스가 바로 종료되어 버림
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
