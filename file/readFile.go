package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	filePath := "/Users/chenyuejun/go/learn/GoHelloWorld/helloworld.go"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileContent))
	openFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	temp := make([]byte, 10)
	index, err := openFile.Read(temp)
	println(string(temp[:index]))
	seek, err := openFile.Seek(100, 1)
	fmt.Println(seek)
	index2, err := openFile.Read(temp)
	fmt.Println(string(temp[:index2]))
	least, err := io.ReadAtLeast(openFile, temp, 8)
	if err != nil {
		panic(err)
	}
	println(least)
	ret, err := openFile.Seek(0, 0)
	println(ret)
	// bufio
	buffReader := bufio.NewReader(openFile)
	peekBytes, err := buffReader.Peek(10)
	println(string(peekBytes))
	err1 := openFile.Close()
	if err1 != nil {
		panic(err1)
	}
}
