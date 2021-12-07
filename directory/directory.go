package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	err := os.Mkdir("testDir", 0644)
	if err != nil {
		return
	}
	defer os.RemoveAll("testDir")

	createEmptyFile := func(fileName string) {
		empty := []byte("")
		check(os.WriteFile(fileName, empty, 0644))
	}
	//createEmptyFile("subdir/abc")
	check(os.MkdirAll("subdir/parent/sub1", 0755))
	createEmptyFile("subdir/parent/sub2")
	createEmptyFile("subdir/parent/sub1/123")
	dir, err := os.ReadDir("subdir/parent")
	for i, entry := range dir {
		println(i)
		fmt.Printf("entry: %v\n", entry)
		println(entry.Type())
		println(entry.Name())
		println(strconv.FormatBool(entry.IsDir()))
	}
	readDir, err := os.ReadDir(".")
	for _, entry := range readDir {
		println(entry.Name(), entry.IsDir())
	}
	os.Chdir("subdir/parent/sub1")
	entries, err := os.ReadDir(".")
	for _, entry := range entries {
		println(entry.Name())
		println(entry.IsDir())
	}
	// 遍历文件夹目录
	//filepath.Walk()
}
