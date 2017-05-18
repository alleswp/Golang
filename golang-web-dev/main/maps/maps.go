package maps

import "fmt"

func Maps() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	delete(m, "k2")
	fmt.Println("map:", m)

	v, ok := m["k1"]
	fmt.Println("ok?:", ok, v)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

	fmt.Println("------------------------------")
	var greeting = make(map[string]string)
	greeting["Wayne"] = "Good Morning"
	greeting["Goober"] = "Hey"
	fmt.Println(greeting)
}
