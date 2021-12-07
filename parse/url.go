package main

import (
	"fmt"
	"net/url"
)

func main() {

	urlStr := "postgres://user:pass@host.com:5432/path?k=v#f"
	parseStr, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parseStr: %v\n", parseStr)
	println(parseStr.Scheme)
	println(parseStr.Path)
	println(parseStr.Fragment)
	println(parseStr.Host)
	fmt.Printf("user: %v\n", parseStr.User)
	println(parseStr.RawPath)
	println(parseStr.RawQuery)
	println(parseStr.Hostname())
	println(parseStr.Port())
	println(parseStr.User.Username())
	password, b := parseStr.User.Password()
	fmt.Println(password, b)
	query, err := url.ParseQuery(parseStr.RawQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println(query)
	println(query["k"][0])

}
