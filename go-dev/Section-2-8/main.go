package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml", "div.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer nf.Close()

	// write nf to file stream
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// write div.gohtml to Stdout
	err = tpl.ExecuteTemplate(os.Stdout, "div.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
