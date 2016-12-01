package main

import "fmt"

func main () {
	c := gen ( 2, 3, 4, 5, 6, 7, 8, 9 )
	out := sq ( c )

	fmt.Println ( <- out )
	fmt.Println ( <- out )
	fmt.Println ( <- out )
	fmt.Println ( <- out )
	fmt.Println ( <- out )
	fmt.Println ( <- out )
	fmt.Println ( <- out )
	fmt.Println ( <- out )
}

func gen ( nums ...int ) chan int {
	input := make ( chan int )
	go func () {
		for _, n := range nums {
			input <- n
		}
		close ( input )
	}()
	return input
}

func sq ( in chan int ) chan int {
	out := make ( chan int )
	go func () {
		for n := range in {
			out <- n * n
		}
		close ( out )
	}()
	return out
}