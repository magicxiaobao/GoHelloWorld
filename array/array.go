package main

import "fmt"

func main() {

	ints := [3]int32{1, 2, 3}
	fmt.Printf("%v", ints)
	fmt.Println(ints)
	var a [3]float64
	fmt.Println(a)
	a[0] = 1.2
	fmt.Println(a)
	fmt.Println(len(a))
	var arrayTwo [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			arrayTwo[i][j] = i + j
		}
	}
	fmt.Println(arrayTwo)
	// array is value types not reference type unlike c c++ java
	intArraySource := [4]int{1, 2, 3, 4}
	intArrayCopyed := intArraySource
	fmt.Println("intArrayCopyed", intArrayCopyed, "intArraySource", intArraySource)
	intArraySource[0] = 99
	fmt.Println("intArrayCopyed", intArrayCopyed, "intArraySource", intArraySource)

	for i, value := range intArraySource {
		fmt.Println("index", i, "value", value)
	}
}
