package udemy

import "fmt"

func Arrays() {
	fmt.Println("------------------------------")
	fmt.Println("Arrays")
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[3])
}