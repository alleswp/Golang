package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/alleswp/Golang/utilities"
	"strings"
)

type Person struct {
	First string
	Last string
	Age int
	secretId int
}

type PersonTagged struct {
	First string
	Last string //`json:"-"`
	Age int `json:"wisdom score"`
}

func main() {

	p1 := Person { "James", "Bond", 20, 007 }
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
	// secretId not exported

	pt := PersonTagged { "Wayne", "Alles", 54 }
	bs, _ = json.Marshal(pt)
	fmt.Println(string(bs))

	unMarshal()
	unMarshalTagged()
	encode()
	decode()
}

func unMarshal() {
	utilities.PrintHeader("Before Marshal")
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte( `{ "First":"James", "Last":"Bond", "Age":20 }` )
	json.Unmarshal(bs, &p1)

	utilities.PrintHeader("After Marshal")

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)

}

func unMarshalTagged() {
	utilities.PrintHeader("Before Tagged Marshal")
	var p1 PersonTagged
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte( `{ "First":"James", "Last":"Bond", "wisdom score":20 }` )
	json.Unmarshal(bs, &p1)

	utilities.PrintHeader("After Tagged Marshal")

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)

}

func encode() {
	utilities.PrintHeader("Encoding")
	p1 := Person { "James", "Bond", 30, 007 }
	json.NewEncoder(os.Stdout).Encode(p1)
}

func decode() {
	utilities.PrintHeader("Decoding")
	var p1 Person
	rdr := strings.NewReader ( `{ "James", "Bond", 30, 007 }` )
	json.NewDecoder(rdr).Decode(&p1)
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}