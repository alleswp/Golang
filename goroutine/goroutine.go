package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go count(i)
		time.Sleep(time.Millisecond * 10000)
	}
}

func count(id int) {

	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)
		time.Sleep(time.Millisecond * 1000)
	}
}
