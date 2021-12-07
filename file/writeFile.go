package main

import (
	"bufio"
	"os"
)

func main() {

	text := "abcdefg"
	err := os.WriteFile("/Users/chenyuejun/go/learn/GoHelloWorld/111.txt", []byte(text), 0644)
	if err != nil {
		return
	}
	createFile, err := os.Create("/Users/chenyuejun/go/learn/GoHelloWorld/222.txt")

	defer createFile.Close()
	writeInt, err := createFile.Write([]byte("hijklmn"))
	println(writeInt)
	_, err1 := createFile.WriteString("123456789")
	if err1 != nil {
		return
	}
	createFile.Sync()

	buffWriter := bufio.NewWriter(createFile)
	buffWriter.Write([]byte("Go is great\n"))
	buffWriter.WriteString("Let's play go")
	buffWriter.Flush()

}
