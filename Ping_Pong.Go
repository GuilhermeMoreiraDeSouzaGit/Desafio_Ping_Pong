package main

import (
	"fmt"
	"time"
)

func Bate(c chan string) {
	for i := 0; ; i++ {
		c <- "Ping"
	}
}

func Rebate(c chan string) {
	for i := 0; ; i++ {
		c <- "Pong"
	}
}
func Mesa(c chan string) {
	for {
		onomatopeia := <-c
		fmt.Println(onomatopeia)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go Bate(c)
	go Mesa(c)
	go Rebate(c)

	var entrada string
	fmt.Scanln(&entrada)
}
