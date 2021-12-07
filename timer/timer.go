package main

import (
	"fmt"
	"time"
)

func main() {

	date1 := time.Date(2021, 11, 1, 1, 1, 1, 1, time.UTC)
	duration := date1.Sub(time.Now())
	fmt.Println(duration.Hours())
	date2 := date1.Add(-duration)
	fmt.Println(date2)
	timer1 := time.NewTimer(2 * time.Second)
	l := <-timer1.C
	fmt.Println("timer1 fired", l)

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()
	time.Sleep(2 * time.Second)
	stopped := timer2.Stop()
	if stopped {
		fmt.Println("timer2 stopped")
	}

}
