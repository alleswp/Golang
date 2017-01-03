package main

import (
	"fmt"
	"github.com/GoesToEleven/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
)

var x = 42

func wrapper() func() int {
	x :=0
	return func() int {
		x++
		return x
	}
}

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Println()


	fmt.Printf("Value: %v\n Type: %T\n\n", a, a)
	fmt.Printf("Value: %v\n Type: %T\n\n", b, b)
	fmt.Printf("Value: %v\n Type: %T\n\n", c, c)
	fmt.Printf("Value: %v\n Type: %T\n\n", d, d)
	foo()

	fmt.Println(vis.MyName)
	vis.PrintVar()

	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())


}

func foo() {
	fmt.Println(x)
}