package main

import "fmt"

func main() {
	fmt.Println("Hello Data Structures")

	//array
	var x [58]string
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	for i := 65; i <= 122; i++ {
		x[i - 65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])


	//slice
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	mySliceOfStrings := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(mySliceOfStrings)
	fmt.Println(mySliceOfStrings[2 : 4])
	fmt.Println(mySliceOfStrings[2])
	fmt.Println("myString"[2])

	//sliceViaMake := make([]int, 25, 50)
	//sliceViaNew := new([50]int) [0:25]

	mySlice2 := make([]int, 0, 5)

	fmt.Println("-----------------------------")
	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
	fmt.Println("-----------------------------")

	for i := 0; i < 21; i++ {
		mySlice2 = append(mySlice2, i)
		fmt.Println("Len: ", len(mySlice2), "Capacity: ", cap(mySlice2), "Value: ", mySlice2[i])
	}

	//2 dimensional slice
	records := make([][]string, 0, 5)

	student1 := make([]string, 4)
	student1[0] = "Kyle"
	student1[1] = "Booth"
	student1[2] = "19"
	records = append(records, student1)

	student2 := make([]string, 4)
	student2[0] = "Eric"
	student2[1] = "Clash"
	student2[2] = "18"
	records = append(records, student2)

	fmt.Println(records)


	mp := make(map[string]int)
	mp["k1"] = 50
	mp["k2"] = 60
	fmt.Println(mp)

	mp2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(mp2)

	v, ok := mp["k2"]
	fmt.Println("ok?: ", ok, v)
	fmt.Println(ok)

	var myGreeting = map[string]string{}
	fmt.Printf("myGreeting is of type: %T\n", myGreeting)
	myGreeting["Wayne"] = "Good Morning"
	myGreeting["Jose"] = "Buenos Dias"
	fmt.Printf("myGreeting is of type: %T\n", myGreeting)
	fmt.Printf("myGreeting: %v\n", myGreeting)

	myGreeting2 := map[int]string {
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	for key, val := range myGreeting2 {
		fmt.Println(key, " - ", val)
	}

	fmt.Println(myGreeting2)

	if val, exists := myGreeting2[2]; exists {
		fmt.Println(val)
		fmt.Println(exists)
		delete(myGreeting2, 2)
	}
	fmt.Printf("After delete, myGreeting2: %v\n", myGreeting2)



}