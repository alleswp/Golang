package main

import "fmt"

func slices2() {
	records := make([][]string, 0)

	std1 := make([]string, 4)
	std1[0] = "Foster"
	std1[1] = "Nathan"
	std1[2] = "100.00"
	std1[3] = "74.00"

	records = append(records, std1)

	std2 := make([]string, 4)
	std2[0] = "Gomez"
	std2[1] = "Lisa"
	std2[2] = "92.00"
	std2[3] = "96.00"

	records = append(records, std2)

	fmt.Println(records)
}
