package main

import "fmt"

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

}