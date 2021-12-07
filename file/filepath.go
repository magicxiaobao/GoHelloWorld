package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	p := filepath.Join("/Users", "/chenyuejun", "/Downloads", "业务监控.jpg")
	fmt.Println(p)
	dir := filepath.Dir("/Users/chenyuejun/Downloads")
	fmt.Println(dir)
	base := filepath.Base("/Users/chenyuejun/Downloads")
	fmt.Println("base: ", base)
	isAbsPath := filepath.IsAbs("/Users")
	fmt.Println(isAbsPath)
	suffix := filepath.Ext("业务监控.jpg")
	fmt.Println(suffix)
	fileName := strings.TrimSuffix("业务监控.jpg", ".jpg")
	fmt.Println(fileName)
	rel, err := filepath.Rel("/Users/chenyuejun", "/Users/chenyuejun/Downloads/业务监控.jpg")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
