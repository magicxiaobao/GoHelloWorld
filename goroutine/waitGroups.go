package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go work(i, &wg)
	}
	wg.Wait()
}

func work(i int, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println("begin to work: ", i)
	time.Sleep(time.Second)
	fmt.Println("success end to work: ", i)
}
