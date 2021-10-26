package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	jobs := make(chan int, 5)
	results := make(chan string, 5)
	for i := 1; i <= 3; i++ {
		go doWork(i, jobs, results)
	}
	for i := 0; i < 5; i++ {
		jobs <- i + 1
	}
	close(jobs)
	for i := 0; i < 5; i++ {
		<-results
	}
	fmt.Println("all works are done")
}

func doWork(i int, jobs <-chan int, results chan<- string) {

	for job := range jobs {
		fmt.Printf("begin to execute job, worker: %d, job: %v\n", i, job)
		time.Sleep(time.Second)
		results <- strconv.Itoa(i * 2)
		fmt.Printf("success to end execute job, worker: %d, job: %v\n", i, job)
	}
}
