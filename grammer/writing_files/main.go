package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// 파일 디스크립터를 통해 운영 체제의 파일 시스템 버퍼에 쓰인 데이터를 디스크에 기록합니다.
	// 운영 체제 수준에서 데이터 일관성을 보장.
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// bufio.Writer의 메모리 버퍼에 저장된 데이터를 기본(writer)에 기록합니다.
	// 애플리케이션 수준에서 버퍼링된 데이터를 출력 스트림으로 전송.
	w.Flush()

	// Sync()와 Flush()는 모두 데이터 기록을 보장하지만,
	// 서로 다른 레벨에서 작동하며 상호 보완적으로 사용됩니다.
	// 일반적으로 bufio.Writer를 사용할 때는 Flush()를 사용하여 버퍼를 비우고,
	// 파일에 직접 쓰기 작업을 할 때는 Sync()를 사용하여 데이터를 디스크에 기록합니다.

}
