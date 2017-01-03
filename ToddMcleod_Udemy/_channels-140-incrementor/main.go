package main

import "fmt"

func main() {
	incFoo := incrementor("Foo: ")
	incBar := incrementor("Bar: ")
	pullFoo := puller(incFoo)
	pullBar := puller(incBar)
	fmt.Println("Final Counter: ", <-pullFoo+<-pullBar)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- i
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
