package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var readCount, writeCount int64
	var data = make(map[int]int)
	mutex := &sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += data[key]
				mutex.Unlock()
				atomic.AddInt64(&readCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				data[key] = value
				mutex.Unlock()
				atomic.AddInt64(&writeCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	readCountFinal := atomic.LoadInt64(&readCount)
	fmt.Println("readCount", readCountFinal)
	writeCountFinal := atomic.LoadInt64(&writeCount)
	fmt.Println("writeCount", writeCountFinal)
	fmt.Println("data", data)
}
