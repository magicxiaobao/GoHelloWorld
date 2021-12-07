package main

import (
	"fmt"
)

func funcReturnMultiResult() (int, int, int) {

	return 1, 2, 3
}

func funcReturnError(price int) (int, error) {

	return price * 2, nil
}

func funcReturnNamedResult(price int) (priceMultied int) {

	priceMultied = price * 2
	return
}

func main() {

	r1, r2, r3 := funcReturnMultiResult()
	fmt.Println(r1, r2, r3)
	priceReturned, err := funcReturnError(12)
	if err != nil {
		panic(err)
	}
	fmt.Println(priceReturned)
	returnNamedResult := funcReturnNamedResult(15)
	fmt.Println(returnNamedResult)
}
