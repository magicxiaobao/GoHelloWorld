package main

import "fmt"

func main() {

	var pointerOfInt *int
	if pointerOfInt == nil {
		fmt.Println("pointerOfInt is nil")
	}
	var int1 = 3
	fmt.Println("int1 memory address", &int1)
	pointerOfInt = &int1
	fmt.Println("pointerOfInt address", pointerOfInt)
	fmt.Println("pointerOfInt value", *pointerOfInt)
	*pointerOfInt = 30
	fmt.Println("after pointer change source value, int1: ", int1)
	pointerInitializedByNew := new(int)
	fmt.Println("pointerInitializedByNew", pointerInitializedByNew)
	*pointerInitializedByNew = 1000
	fmt.Println("pointerInitializedByNew value", *pointerInitializedByNew)
	var pointerOfInt2 = &int1
	fmt.Println("pointerOfInt1 equals pointerOfInt2", pointerOfInt == pointerOfInt2)

}
