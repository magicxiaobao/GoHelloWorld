package main

import (
	"fmt"
	"math"
)

func main() {

	// 无确定类型的constant
	const a = 1
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	i := a / 3
	fmt.Println(i)
	sin := math.Sin(a)
	fmt.Println(sin)
	// 指定类型的constant
	const b int64 = 1
	fmt.Printf("%T\n", b)
	var bint32 int32 = int32(b)
	fmt.Printf("%T\n", bint32)
	const c = 1 + 1.5
	fmt.Printf("%T\n", c)
	const d = 'a' + 1
	fmt.Printf("%T\n", d)
	const e = 'C' + 1
	fmt.Printf("%T\n", e)

}
