package main

import "fmt"

func main() {

	sliceInt1 := make([]int, 5)
	fmt.Println(sliceInt1, len(sliceInt1))
	for i := 0; i < 10; i++ {
		sliceInt1 = append(sliceInt1, i)
	}
	fmt.Println(sliceInt1)
	sliceInt1[0] = 100
	fmt.Println(sliceInt1)
	sliceInt2 := make([]int, 3)
	copy(sliceInt2, sliceInt1)
	fmt.Println(sliceInt2)
	sliceInt2[1] = 90
	sliceInt2[2] = 80
	fmt.Println("after copy, sliceInt2: ", sliceInt2, " sliceInt1: ", sliceInt1)
	slice2Part := sliceInt2[1:2]
	fmt.Println(slice2Part)
	slice2Part[0] = 99
	fmt.Println(sliceInt2)
	slice2Part2 := sliceInt2[:len(sliceInt2)-1]
	fmt.Println(slice2Part2)
	slice2Part3 := sliceInt2[2:]
	fmt.Println(slice2Part3)
	stringSliceInitialized := []string{"a", "b", "c"}
	fmt.Println(stringSliceInitialized)
	sliceInt3 := []int{1, 2, 3, 4, 5}
	sliceInt3Part := sliceInt3[:5]
	fmt.Println(sliceInt3Part)
	sliceOfSliceInt3 := sliceInt3Part[2:4]
	fmt.Println("before change sliceOfSliceInt3: ", sliceOfSliceInt3, "sliceInt3: ", sliceInt3)
	sliceOfSliceInt3[0] = 999
	fmt.Println("after change sliceOfSliceInt3: ", sliceOfSliceInt3, "sliceInt3: ", sliceInt3)
	intArray := []int{1, 2, 3, 4, 5}
	sliceOfIntArray := intArray[1:3]
	fmt.Println("length: ", len(sliceOfIntArray), "capacity: ", cap(sliceOfIntArray)) // capacity is from underlying array index to end
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original Slice")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[1:5]
	fmt.Println("\nAfter slicing from index 1 to 5")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[:8]
	fmt.Println("\nAfter extending the length")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[2:]
	fmt.Println("\nAfter dropping the first two elements")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	sliceInt4 := make([]int, 5, 10)
	fmt.Println("sliceInt4: ", len(sliceInt4), "capacity: ", cap(sliceInt4))
	var nilSlice []int
	fmt.Printf("nilSlice: %v,lengh: %d,capacity: %d\n", nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nilSlice is nil")
	}
	// append slice 原slice的数组不够 会创建新的数组
	slice1 := make([]string, 3, 10)
	copy(slice1, []string{"C", "C++", "Java"})
	slice2 := append(slice1, "Python", "Ruby", "Go")
	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))
	slice1[0] = "C#"
	fmt.Println("\nslice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	slice11 := []int{1, 2, 3, 4, 5}
	slice22 := []int{6, 7, 8, 9, 10}
	slice33 := append(slice11, slice22...)
	fmt.Printf("slice3: %v, slice33, len: %d, capacity: %d\n", slice33, len(slice33), cap(slice33))

}
