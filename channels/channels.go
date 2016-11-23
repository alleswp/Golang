package main

import (
	"fmt"
	"time"
	"strconv"
)

var pizzaNum = 0
var pizzaName = ""

func main() {
	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
		time.Sleep(time.Millisecond * 3000)
	}

}

func makeDough(stringChan chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("Make dough and send ", pizzaName, " for sauce")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	pizza := <- stringChan


	fmt.Println("Add Sauce and send ", pizza, " for toppings")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChan chan string) {
	pizza := <- stringChan
	fmt.Println("Add toppings to ", pizza, " and ship")
	time.Sleep(time.Millisecond * 10)
}
