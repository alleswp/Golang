package main

import "fmt"

type foo int

type Person struct {
	name string
	age int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (p *Person) SetAge(newAge int) {
	p.age = newAge
}

func (dz DoubleZero) Greeting() {
	fmt.Println("I'm not who you think I am")
}

func main() {

	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	p1 := Person{
		name: "Ian Fleming",
		age: 20,
	}

	p2 := DoubleZero{
		Person: Person{
			name: "James Bond",
			age: 29,
		},
		LicenseToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()

	p3 := &Person{"Wayne", 54}
	fmt.Printf("%T\n", p3)
	fmt.Println(p3.name)
	fmt.Println(p3.age)
	p3.SetAge(18)
	fmt.Println(p3.age)


}
