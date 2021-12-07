package main

import "fmt"

func main() {

	emptyMap := make(map[string]int, 10)
	fmt.Println("emptyMap", emptyMap, "size", len(emptyMap))
	emptyMap["a"] = 1
	fmt.Println(emptyMap)
	fmt.Println(emptyMap["a"], emptyMap["b"])
	stringMap := make(map[string]string)
	stringMap["a"] = "A"
	fmt.Println(stringMap["a"], stringMap["b"])
	isNullValue, isExist := stringMap["b"]
	fmt.Printf("%v\n", isExist)
	if isNullValue == "" {
		fmt.Println("isNullValue is blank")
	}
	fmt.Printf("%v", isNullValue)
	delete(stringMap, "b")
	fmt.Println("after deleted no exist key", stringMap)
	delete(stringMap, "a")
	fmt.Println("after delete key", stringMap["a"])
	mapInitialized := map[string]int{"a": 1}
	fmt.Println("mapInitialized", mapInitialized)
}
