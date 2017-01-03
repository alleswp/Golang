package main

import (
	"fmt"
)

func main() {
	inc_foo := incrementor("Foo:")
	inc_bar := incrementor("Bar:")
	pull_foo := puller(inc_foo)
	pull_bar := puller(inc_bar)
	fmt.Println("Final Counter:", <- pull_foo, + <- pull_bar)

}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}