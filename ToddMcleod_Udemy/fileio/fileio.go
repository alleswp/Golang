package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"

)

func main() {

	writeFile()
	readFile()


}

func writeFile() {

	if file, err := os.Create("Sample.txt"); err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()
		file.WriteString("This is some random text")
	}
}

func readFile() {

	if stream, err := ioutil.ReadFile("Sample.txt"); err != nil {
		log.Fatal(err)
	} else {
		readString := string(stream)
		fmt.Println(readString)
	}
}