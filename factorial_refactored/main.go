package main

import "fmt"

func main() {
	//set up pipeline and consume the output
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)

	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
			//fmt.Printf("from gen: %d \n", n)
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
			//fmt.Printf("from sq: %d \n", n*n)
		}
		close(out)
	}()
	return out
}
