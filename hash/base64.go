package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	base64Result := base64.StdEncoding.EncodeToString([]byte("abcdefg"))
	println(base64Result)
	decodeBytes, err := base64.StdEncoding.DecodeString(base64Result)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decodeBytes))
	urlEncoded := base64.URLEncoding.EncodeToString([]byte("abcdeft"))
	println(urlEncoded)
	decodeString, err := base64.URLEncoding.DecodeString(urlEncoded)
	if err != nil {
		panic(err)
	}
	println(string(decodeString))
}
