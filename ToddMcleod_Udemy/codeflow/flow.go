package main

import "fmt"

func main() {

	for i := 1; i <= 3; i++ {
		fmt.Printf("Outer loop value: %d\n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("Inner loop value: %d\n", j)

		}
	}

	f := 1
	for f <= 3 {
		fmt.Printf("For Loop: %d\n", f)
		f++
	}

	g := 1

	//Print only even numbers between 1 and 10
	for {
		if g > 10 {
			break
		}
		if g % 2 == 0 {
			fmt.Printf("Even Number: %d\n", g)
		}
		g++

	}

	ch := 'a'
	fmt.Println(string(ch))
	fmt.Printf("%T \n", ch)




}
