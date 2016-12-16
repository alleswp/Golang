package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go beginCount(c)
	for n := range c {
		fmt.Println(n)
	}
}

func beginCount(c chan int) {
	for i := 1; i <= 20; i++ {
		c <- i
	}
	close(c)
}
