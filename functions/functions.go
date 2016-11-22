package main

import "fmt"

var x int

func increment() int {
	x++
	return x
}

func wrapper() func() int {
	var x int = 10
	return func() int {
		x++
		return x
	}
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
		xs = append(xs, n)
		}
	}
return xs
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}

func main() {
	fmt.Println("Hello Functions")
	n := average(43, 259, 87, 12, 45, 517)
	fmt.Println(n)
	data := []float64{ 24, 65, 47, 98, 43 }
	m := average2(data)
	fmt.Println(m)
	fmt.Printf("%T\n", average)
	fmt.Printf("%T\n", average2)
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	incrementer := wrapper ()
	fmt.Println(incrementer())
	fmt.Println(incrementer())
	fmt.Println(incrementer())

	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 2
	})
	fmt.Println(xs)
	fmt.Println(factorial(4))

	age := 44
	fmt.Println(&age)

	changeMe(&age)
	fmt.Println(&age)
	fmt.Println(age)

	mm := make([]string, 1, 25)
	fmt.Println(mm)
	changeMe2(mm)
	fmt.Println(mm)

	mp := make(map[string]int)
	changeMap(mp)
	fmt.Println(mp["Todd"])

	func() {
		fmt.Println("I'm Driving")
	}()

}

func changeMap(z map[string]int) {
	z["Todd"] = 44
}

func changeMe2(z []string) {
	z[0] = "Todd"
	fmt.Println(z)
}

func changeMe(z *int) {
	fmt.Println("changeMe called")
	fmt.Println(z)
	fmt.Println(*z)
	*z = 24
	fmt.Println(z)
	fmt.Println(*z)

}


func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	fmt.Println(len(sf))
	return total / float64(len(sf))

}

func average2(sf []float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	fmt.Println(len(sf))
	return total / float64(len(sf))
}

func makeGreeter() func() string {
	return func() string {
		return "Hello World"
	}
}