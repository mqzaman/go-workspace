package main

import (
	"fmt"
	"time"
)

func Server1(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "From Server1"
}
func Server2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "From Server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go Server1(output1)
	go Server2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)

	}
}
