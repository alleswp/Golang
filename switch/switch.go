package main

import "fmt"

var num int = 0



func main() {
	fmt.Println("Enter a number from 1 to 5:")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("You Entered One")
		main()
	case 2:
		fmt.Println("You Entered Two")
		main()
	case 3:
		fmt.Println("You Entered Three")
		main()
	case 4:
		fmt.Println("You Entered Four")
		main()
	case 5:
		fmt.Println("You Entered Five")
		main()
	default:
		switch {
		case num > 5 || num < 0:
			fmt.Println("You entered a number out of range. Try again")
			main()
		case num == 0:
			//quit
		}
	}
}
