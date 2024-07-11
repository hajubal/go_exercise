package main

import (
	"embed"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	println("fileString: ", fileString)
	println("fileByte: ", string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	println("file1: ", string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	println("file2: ", string(content2))
}
