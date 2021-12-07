package main

import (
	"fmt"
	"strconv"
)

func main() {

	parseInt1, err := strconv.ParseInt("123", 0, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(parseInt1)
	parseBool, err := strconv.ParseBool("1")
	if err != nil {
		panic(err)
	}
	fmt.Println(parseBool)
	float, err := strconv.ParseFloat("2.323232", 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(float)
	parseUint, err := strconv.ParseUint("2121", 0, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(parseUint)
	atoi, err := strconv.Atoi("3232")
	if err != nil {
		panic(err)
	}
	fmt.Println(atoi)
	itoa := strconv.Itoa(232)
	fmt.Println(itoa)
	/*i, err := strconv.Atoi("fewe")
	if err != nil {
		panic(err)
	}
	fmt.Println(i)*/
}
