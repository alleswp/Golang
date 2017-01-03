package main

import (
	"fmt"
)

func main() {

	c := make ( chan int )

	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println ( n )
	}
}
