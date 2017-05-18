package udemy

import "fmt"

func Slices() {
	fmt.Println("------------------------------")
	fmt.Println("Slices")
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	fmt.Println("MyInt: "[2])
	fmt.Println(string("MyInt: "[2]))
	fmt.Println("------------------------------")
	// 2 ways to make a slice
	s1 := make([]int, 5, 10)
	s2 := new([10]int)[0:5]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("------------------------------")

	s3 := make([]int, 0, 3)
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
	fmt.Println("------------------------------")

	for i := 0; i < 12; i++ {
		s3 = append(s3, i)
		fmt.Println("Len: ", len(s3), "Capacity: ", cap(s3), "Value: ", s3[i])
	}
	fmt.Println("------------------------------")

	greeting := []string{
		"Good Morning",
		"Bonjour",
		"Aloha",
	}
	for i, entry := range greeting {
		fmt.Println(i, entry)
	}
	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
	fmt.Println("------------------------------")

	s4 := []string{"monday", "tuesday"}
	s5 := []string{"wednesday", "thursday", "friday"}
	combined := append(s4[:2], s5...)
	fmt.Println(combined)
}
