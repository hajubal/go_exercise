package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	log.Println("standard logger")
	// 2024/04/26 21:53:33 standard logger

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")
	// 2024/04/26 21:55:37.774287 with micro

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")
	// 2024/04/26 21:55:37 main.go:21: with file/line

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")
	// my:2024/04/26 21:55:37 from mylog

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")
	// ohmy:2024/04/26 21:55:37 from mylog

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String())
	// from buflog:buf:2024/04/26 21:55:37 hello

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	// {"time":"2024-04-26T21:55:37.774381+09:00","level":"INFO","msg":"hi there"}

	myslog.Info("hello again", "key", "val", "age", 25)
	// {"time":"2024-04-26T21:55:37.774396+09:00","level":"INFO","msg":"hello again","key":"val","age":25}
}
