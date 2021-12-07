package main

import (
	"fmt"
	"sort"
)

func main() {

	intsSlice := []int{1, 2, 6, 3}
	sort.Ints(intsSlice)
	fmt.Println(intsSlice)
	stringsSlice := []string{"a", "C", "g", "e"}
	sort.Strings(stringsSlice)
	fmt.Println(stringsSlice)
	stringSorted := sort.StringsAreSorted(stringsSlice)
	fmt.Println(stringSorted)

	students := []student{{"guoguo", 9}, {"suisui", 37}}
	sort.Sort(studentByAge(students))
	fmt.Println(students)

}

type student struct {
	name string
	age  int
}

type studentByAge []student

func (s studentByAge) Len() int {
	return len(s)
}

func (s studentByAge) Less(i, j int) bool {
	return s[i].age > s[j].age
}

func (s studentByAge) Swap(i, j int) {
	var t = s[i]
	s[i] = s[j]
	s[j] = t
}
