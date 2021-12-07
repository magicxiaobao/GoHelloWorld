package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	base64Encoded := base64.StdEncoding.EncodeToString([]byte("aaaa"))
	fmt.Println(base64Encoded)
	decodeString, err := base64.StdEncoding.DecodeString(base64Encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decodeString))
}
