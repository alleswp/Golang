package main

import (
	"fmt"
	//"github.com/therecipe/qt"
)

func main() {
	num, err := fmt.Println("Hello World")
	fmt.Printf("Number: %d\n", num)
	fmt.Println(err)
}
