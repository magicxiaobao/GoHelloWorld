package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func printSlice(slice []string) error {
	for k, v := range slice {
		fmt.Printf("key:%v, value:%v\n", k, v)
	}
	return nil
}

func main() {
	fmt.Println("go come on")
	fmt.Println(1 + 5)
	fmt.Printf("%v, %T\n", math.Sqrt(9), math.Sqrt(9))
	var x = "Xiaobao"
	fmt.Printf("%v, %T\n", x, x)
	var y = 360
	fmt.Printf("%v, %T\n", y, y)
	var z = 20.88
	fmt.Printf("%v, %T\n", z, z)
	fmt.Printf("name is %v, age is %v\n", x, y)
	fmt.Printf("the player name is %v, hobby is %v\n", x, "swimming")
	fmt.Println(9, "guoguo", "singing", z)
	fmt.Printf("%v\n", "A" == "\x41")
	var a = "SuiSui \"小朋友\""
	var b = "mongo"
	var s = fmt.Sprintf("name is %v, favorite fruit is %v\n", a, b)
	fmt.Println(s)
	var rawString = `ni hao ma
 many lines
		tab too\n`
	fmt.Println(rawString)
	var abc = "abc"
	fmt.Printf("%#v\n", abc[0])
	fmt.Printf("%T\n", abc[1])
	fmt.Printf("%#v\n", abc[2])
	fmt.Println("a\nb\n")
	fmt.Printf("%v\n", len("abc"))
	fmt.Printf("%v\n", len("你好吗?"))
	fmt.Printf("%v\n", len("你"))
	fmt.Printf("%v\n", len("?"))
	// substring
	var sourceString = "0123456"
	fmt.Printf("result %#v, sourceString %#v\n", sourceString[2:4], sourceString)
	var stringNoAsc2 = "👋👋👋"
	fmt.Printf("%#v, len is %v,\njoinString is %v\n", stringNoAsc2, len(stringNoAsc2), stringNoAsc2+"拍掌")
	fmt.Printf("%#v\n", stringNoAsc2[2:3])
	fmt.Println("ab" == "ab")
	fmt.Println(strings.Contains("abcdeft", "abc"))
	fmt.Println(strings.Join([]string{"A", "B", "C"}, "&"))
	fmt.Printf("%T\n", strings.Split("x y z", " "))
	fmt.Printf("%d\n", "abc"[0])
	fmt.Printf("%q\n", "abc"[0]) // print as byte
	fmt.Printf("%T\n", 1)
	var int1 = 12
	var float1 = float32(int1)
	fmt.Printf("%T\n", float1)
	fmt.Println(reflect.TypeOf(float1))
	var c1 = '😂'
	fmt.Printf("%q %d\n", c1, c1)
	var a2, b2, c2 int
	var x2 int
	var y2 float32
	var z2 string
	fmt.Printf("a2: %#v, b2: %#v, c2: %#v, x2: %#v, y2: %#v, z2: %#v\n", a2, b2, c2, x2, y2, z2)
	var a3, b3 = 60, 100
	fmt.Println(a3, b3)
	a4 := 999
	fmt.Println(a4)
	var (
		a5 int
		b5 string
		c5 = "multiVariableDeclaration"
	)
	fmt.Println(a5, b5, c5)
	fmt.Printf("%#v\n", b5)
	var isHungry = true
	fmt.Println(isHungry)
	// zero value nil
	var sl []int
	fmt.Printf("%#v\n", sl)
	fmt.Printf("%v\n", sl)
	// const
	const const1 = 10
	const const2 = "abc"
	fmt.Printf("%v %v\n", const1, const2)
	// if else
	if const1 > 10 {
		fmt.Printf("const1 > 10\n")
	} else if const1 == 10 {
		fmt.Printf("const1 == 10\n")
	} else {
		fmt.Printf("const1 < 10\n")
	}
	if k := 8; const1 > k {
		fmt.Printf("const1 > k\n")
	}
	var j = 5
	switch j {
	case 5, 0:
		fmt.Println("j == 5 0 or j == 0")
		fallthrough // 贯穿
	case 3, 4:
		fmt.Println("j == 3 or j == 4")
	}
	switch {
	case const1 > 10:
		fmt.Printf("const1 > 10\n")
	case const1 == 10:
		fmt.Printf("const1 == 10\n")
		fallthrough
	default:
		fmt.Printf("const1 < 10\n")
	}
	i := 10
	for {
		fmt.Printf("i:%v\n", i)
		i++
		if i == 15 {
			break
		}
	}
	for i := 5; i < 10; i++ {
		fmt.Printf("i=%v\n", i)
	}
	for i < 20 {
		fmt.Printf("i=%v\n", i)
		i++
	}
	for i := 6; i < 12; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("i=%v\n", i)
	}
	var intArray1 [10]int
	intArray1[0] = 1
	intArray1[1] = 3
	intArray1[5] = 6
	fmt.Printf("intArray1%v, length:%v, capacity:%v\n", intArray1, len(intArray1), cap(intArray1))
	var intArray2 = [3]int{1, 2, 3}
	fmt.Printf("intArray2%v\n", intArray2)
	var intArray3 = [...]string{"a", "b", "c"}
	intArray3[2] = "CC"
	fmt.Printf("intArray3%v\n", intArray3)
	for i, x := range intArray1 {
		fmt.Printf("i:%v,v:%v\n", i, x)
	}
	for _, x := range intArray2 {
		fmt.Printf("item:%v\n", x)
	}
	// slice
	var array3 = []int{1, 2, 4}
	var slice3 = array3[:]
	fmt.Printf("type:%T,value:%v\n", slice3, slice3)
	slice3[2] = 5
	fmt.Printf("array3:%v, slice3:%v\n", array3, slice3)
	var slice4 []string
	fmt.Printf("slice4 type:%T\n", slice4)
	var slice5 = []string{"A", "b", "C"}
	fmt.Printf("slice5:%#v,%v\n", slice5, slice5)

	var bookList = []string{"雪山飞狐.html", "笑傲江湖.pdf"}
	printSlice(bookList)
	var slice6 = make([]string, 10)
	fmt.Printf("%#v\n", slice6)
	var slice7 = make([]int, 5, 10)
	fmt.Printf("len:%v,cap:%v,%#v\n", len(slice7), cap(slice7), slice7)
	slice7[0] = 3
	slice7[1] = 4
	slice7[2] = 5
	fmt.Printf("len:%v,cap:%v,%#v\n", len(slice7), cap(slice7), slice7)
	var slice8 = slice7[1:]
	fmt.Printf("sclice8%v\n", slice8)
	slice8[0] = 99
	fmt.Printf("sclice8%v,slice7:%v\n", slice8, slice7)
	// slice append
	var slice9 = []int{1, 2, 3}
	var slice9AfterAppend = append(slice9, 88, 99)
	slice9AfterAppend[0] = 100
	fmt.Printf("slice9AfterAppend%v, slice9:%v\n", slice9AfterAppend, slice9)
	var slice10 = make([]int, 4, 10)
	slice10[0] = 1
	slice10[1] = 2
	slice10[2] = 3
	slice10[3] = 4
	var slice10AfterAppend = append(slice10, 88, 99)
	slice10AfterAppend[0] = 100
	fmt.Printf("slice10AfterAppend%v, slice10:%v\n", slice10AfterAppend, slice10)
	var x21 = []int{0, 1, 2, 3, 4, 5}
	var x22 = x21[:3] // [0 1 2]
	var x23 = append(x22, 22)
	var x24 = append(x22, 23, 24)
	// x21 is changed
	fmt.Println(x21) // [0 1 2 22 4 5]
	fmt.Println(x22)
	fmt.Println(x23) // [0 1 2 22]
	fmt.Println(x24)
	var slice11 = []int{1, 2, 3}
	var slice12 = []int{4, 5, 6}
	var slice13 = append(slice11, slice12...)
	slice11[1] = 99
	fmt.Printf("slice11: %v, slice13: %v\n", slice11, slice13)
	var slice14 = make([]string, 4, 10)
	slice14[0] = "1"
	var slice15 = append(slice14, []string{"GG"}...)
	fmt.Printf("slice14: %v, slice15: %v\n", slice14, slice15)
	// use slice to delete element
	var slice16 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice17 = append(slice16[:3], slice16[5:]...)
	fmt.Printf("slice17: %v\n", slice17)
	// copy slice
	var source1 = []int{1, 2, 3, 4, 5}
	var dest1 = []int{10, 20}
	var copy1 = copy(dest1, source1)
	fmt.Printf("copy1: %v, source1: %v, dest1: %v\n", copy1, source1, dest1)
	var source2 = []int{1, 2}
	var dest2 = []int{10, 20, 30, 40, 50, 60}
	var copy2 = copy(dest2, source2)
	fmt.Printf("copy2:%v, source2: %v, dest2 %v\n", copy2, source2, dest2)
	// clear slice
	var slice18 = []int{1, 2, 3, 4, 5}
	slice18 = nil
	var slice19 = append(slice18, []int{10, 20}...)
	fmt.Printf("slice18: %v, slice19: %v\n", slice18, slice19)
	slice19 = slice19[0:0]
	fmt.Println("slice19:", slice19)
	// nested slice
	var nestedSlice = [][]int{{1, 2}, {10, 20}}
	fmt.Println(nestedSlice)
	var nestedStringSlice = make([][]string, 3, 10)
	nestedStringSlice[0] = []string{"a", "b"}
	nestedStringSlice[1] = []string{"A", "B"}
	var string01 = nestedStringSlice[0][1]
	fmt.Println(nestedStringSlice, string01)
	nestedStringSlice[1][0] = "GG"
	fmt.Println(nestedStringSlice)
	// Join String slice
	var stringSlice = []string{"a", "b", "c"}
	var joinStringResult = strings.Join(stringSlice, ";")
	fmt.Println(joinStringResult)
	for k, v := range stringSlice {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
	for _, v := range stringSlice {
		fmt.Println(v)
	}
	// map
	// declare map
	var map1 map[string]int
	fmt.Printf("%T\n", map1)
	// create
	var map2 = map[string]int{"a": 1, "A": 10}
	fmt.Printf("map2: %v\n", map2)
	// create by make
	var map3 = make(map[string]int)
	map3["xiaobao"] = 33
	map3["guoguo"] = 9
	fmt.Printf("map3: %v\n", map3)
	var v1, e1 = map3["suisui"]
	fmt.Printf("v: %v, e: %v\n", v1, e1)
	var v2 = map3["xiaobao"]
	fmt.Printf("v2: %v\n", v2)
	map3["xiaobao"] = 34
	map3["suisui"] = 36
	fmt.Printf("xiaobao: %v, suisui: %v\n", map3["xiaobao"], map3["suisui"])
	delete(map3, "lanlan")
	fmt.Println(map3)
	delete(map3, "suisui")
	fmt.Println(map3)
	for k, _ := range map3 {
		fmt.Println(k)
	}
	// struct
	var s1 Student = Student{"aa", 1}
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s1: %+v\n", s1)
	var s2 = Student{name: "bb", age: 2}
	fmt.Printf("s2: %v\n", s2)
	var s3 = Student{}
	fmt.Printf("s3: %v\n", s3)
	s3.name = "xxxxyy"
	fmt.Printf("s3.name: %#v\n", s3.name)
	var s4 = Student{"gg", 100}
	fmt.Printf("s4: %v\n", s4)
	// function
	var plusResult = plus(2, 4)
	fmt.Printf("plusResult: %v\n", plusResult)
	var fun1Result = func1("a", "b", 5)
	fmt.Printf("fun1Result: %v\n", fun1Result)
	var fp1, fp2, fp3 = func2()
	fmt.Println(fp1, fp2, fp3)
	var fun3Result = func3()
	fmt.Printf("fun3Result: %v\n", fun3Result)
	var func4Result = func(x int) int {
		return x + 5
	}(2)
	fmt.Printf("func4Result: %v\n", func4Result)
	func5 := func(i int) int {
		return i + 5
	}
	fmt.Println(func5(20))
	var func6Result = func6(plus, 5)
	fmt.Printf("func6Result: %v\n", func6Result)
	// return func
	returnFuncResult := returnFunc(2)(2)
	fmt.Printf("returnFuncResult: %v\n", returnFuncResult)
	var unSpecifiedResult = unSpecifiedArg(1, 2, 3)
	fmt.Printf("unSpecifiedResult: %v\n", unSpecifiedResult)
	var arrayArgResult = arrayArgFunc([]int{2, 3, 4}...)
	fmt.Printf("arrayArgResult: %v\n", arrayArgResult)
	// closure
	var closureResult = funReturnClosure()
	fmt.Printf("clucureResult: %T, closureInvokeResult: %v\n", closureResult, closureResult(5))
	fmt.Printf("clusureInvokeTwice: %v\n", closureResult(5))
	// pointer
	// delcare
	var point1 *int
	var int11 = 1
	point1 = &int11
	fmt.Printf("%v\n", point1)
	var int22 = *point1
	fmt.Printf("int22: %v\n", int22)
	studentOne := Student{"guoguo", 9}
	pOfStudentOne := &studentOne
	pOfStudentOne.age = 100
	fmt.Printf("pointOfStudentOne: %v\n", pOfStudentOne)
	var intParam1 = 2
	var funParamIsPointResult = funParmIsPointer(&intParam1)
	fmt.Printf("intParam1: %v,funParamIsPointResult: %v\n", intParam1, funParamIsPointResult)
	// regexp
	var longText1 = `xiaobao eating apple`
	var re = regexp.MustCompile(`(?i:xiaobao)`)
	var replaceText1 = "guoguo"
	var result1 = re.ReplaceAllLiteralString(longText1, replaceText1)
	fmt.Printf("result1: %v\n", result1)
	// read whole file
	fileText, fileError := ioutil.ReadFile("/Users/chenyuejun/go/learn/gobyexample/README.md")
	if fileError != nil {
		panic(fileError)
	}
	fmt.Printf("fileText: %T\n", fileText)
	fmt.Println(string(fileText))
	// read first x bytes
	var path = "/Users/chenyuejun/go/learn/gobyexample/README.md"
	var bytes1 = getHeadBytes(path, 5)
	fmt.Printf("bytes1: %v\n", bytes1)
	// write to file
	writeText := []byte(`go is beautiful`)
	outPath := "goIsBeautiful.txt"
	err := ioutil.WriteFile(outPath, writeText, 0644)
	if err != nil {
		panic(err)
	}
	// use os lib
	writeText1 := "java is boring"
	outPath1 := "javaIsBoring.txt"
	file, err := os.Create(outPath1)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Printf("fileTypeIs: %T\n", file)
	writeLen1, errW := file.WriteString(writeText1)
	if errW != nil {
		panic(errW)
	}
	fmt.Printf("writeLen1: %v\n", writeLen1)
	// go through file
	var goThroughPath = "/Users/chenyuejun/go/test_go"
	var ff = func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error %v at path %q\n", err, path)
			return err
		}
		fmt.Printf("path: %v\n", path)
		if fileInfo.IsDir() {
			fmt.Printf("is dir.\n")
		} else {
			fmt.Printf("dir: %v\n", filepath.Dir(path))
			fmt.Printf("file name: %v\n", fileInfo.Name())
			fmt.Printf("extenion: %v\n", filepath.Ext(path))
		}
		return nil
	}
	error1ByWalkingFile := filepath.Walk(goThroughPath, ff)
	if error1ByWalkingFile != nil {
		fmt.Printf("error1ByWalkingFile: %v\n", error1ByWalkingFile)
	}
	var fileExistTrue = isFileExist("/Users/chenyuejun/go/test_go")
	fmt.Printf("fileExistTrue: %v\n", fileExistTrue)
	var fileExistFalse = isFileExist("/Users/chenyuejun/go/test_go1")
	fmt.Printf("fileExistFalse: %v\n", fileExistFalse)

	// exec
	var cmdName1 = "ls"
	var changeDir = "/Users/chenyuejun/go"
	errOfChangeDir := os.Chdir(changeDir)
	if errOfChangeDir != nil {
		fmt.Printf("errOfChangeDir: %v\n", errOfChangeDir)
	}
	command1 := exec.Command(cmdName1, "-a", "-l")
	output, err := command1.Output()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("output: %v\n", string(output))
	// get script path
	var currentScriptPath = getCurrentScriptPath()
	fmt.Printf("currentScriptPath: %v\n", currentScriptPath)
	// defer 延迟执行方法，方法里的参数会立刻执行
	var funcDefer = func(x int) int {
		fmt.Println("funcDefer invoked")
		return x + 1
	}
	fmt.Println("print before defer")
	defer funcDefer(getInput())
	fmt.Println("print after defer")
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	rand.Seed(100)
	for index := 0; index < 30; index++ {
		fmt.Printf("%v", rand.Intn(10))
	}
	fmt.Println("")
}


func getInput() int {
	fmt.Println("invoke getInput")
	return 100
}

func getCurrentScriptPath() string {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return path
}

func isFileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		return false
	}
}

func getHeadBytes(path string, n int) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	headBytes := make([]byte, n)
	m, err := file.Read(headBytes)
	if err != nil {
		panic(err)
	}
	return headBytes[:m]
}

func funParmIsPointer(p *int) int {

	*p = 4
	return *p
}

type Student struct {
	name string
	age  int
}

func plus(x int, y int) int {
	return x + y
}
func func1(x, y string, z int) string {
	return x + y + strconv.Itoa(z)
}
func func2() (int, int, int) {
	return 3, 4, 5
}
func func3() (r int) {
	r = 1 + 2
	return r
}
func func6(f func(int, int) int, p int) int {
	return f(p, p)
}
func returnFunc(x int) func(int) int {
	returnFunc := func(y int) int {
		return x + 1 + y
	}
	return returnFunc
}
func unSpecifiedArg(x ...int) []int {
	return x
}
func arrayArgFunc(x ...int) int {
	var total = 0
	for _, v := range x {
		total = total + v
	}
	return total
}
func funReturnClosure() func(int) int {
	var i = 5
	var f = func(v int) int {
		i += v
		return i
	}
	return f
}

