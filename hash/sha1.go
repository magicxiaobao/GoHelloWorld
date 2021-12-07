package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	sha1 := sha1.New()
	_, err := sha1.Write([]byte("abcdefg"))
	if err != nil {
		panic(err)
	}
	sum := sha1.Sum(nil)
	fmt.Printf("result: %x\n", sum)

}
