package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Code    int
	Message string
	data    interface{}
}

type Response2 struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {

	b, _ := json.Marshal(true)
	fmt.Println(string(b))
	one, _ := json.Marshal(1)
	fmt.Println(string(one))
	onePoint2, _ := json.Marshal(1.2)
	fmt.Println(string(onePoint2))
	stringJson, _ := json.Marshal([]string{"a", "b"})
	fmt.Println(string(stringJson))
	mjson, _ := json.Marshal(map[string]int{"a": 1, "b": 2})
	fmt.Println(string(mjson))
	response1 := &Response1{
		0, "", nil,
	}
	r1, _ := json.Marshal(response1)
	fmt.Println(string(r1))
	response2 := &Response2{
		500, "error occurred", nil,
	}
	r2, _ := json.Marshal(response2)
	fmt.Println(string(r2))

	var d map[string]interface{}
	bytes := []byte(`{"age": 33,"name":"xiaobao"}`)
	if err := json.Unmarshal(bytes, &d); err != nil {
		panic(err)
	}
	fmt.Println(d["age"].(float64))
	fmt.Println(d["name"].(string))
	var r22 Response2
	bb := []byte(`{"code":1,"message":"aaa","data":["a","b"]}`)
	if err := json.Unmarshal(bb, &r22); err != nil {
		panic(err)
	}
	fmt.Println(r22, r22.Code, r22.Data)
	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(map[string]int{"a": 111, "b": 222})
	if err != nil {
		panic(err)
	}

}
