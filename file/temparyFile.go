package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	temp, err := os.CreateTemp("", "sample")
	if err != nil {
		panic(err)
	}
	println(temp.Name())
	temp.WriteString("abcd")
	defer os.Remove(temp.Name())
	dirTemp, err := os.MkdirTemp("", "111")
	if err != nil {
		panic(err)
	}
	println(dirTemp)
	defer os.RemoveAll(dirTemp)
	joinFileName := filepath.Join(dirTemp, "file1")
	os.WriteFile(joinFileName, []byte{1, 2}, 0666)
	bytes, err := os.ReadFile(joinFileName)
	fmt.Println(bytes)
}
