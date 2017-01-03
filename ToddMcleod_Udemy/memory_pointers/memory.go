package main

import "fmt"

const metersToYards float64 = 1.09361

func zero(x *int) {
	*x = 0
}

func main() {
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)

	var meters float64
	fmt.Println("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")

	var b *int = &a

	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)

	x := 5
	zero(&x)
	fmt.Println(x)

}
