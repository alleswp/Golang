package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl-map.gohtml"))
}

func main() {

	// sages := []string{"Gandhi", "MLK", "Buddha", "Jesus",
	// 	"Muhammad"}

	// err := tpl.Execute(os.Stdout, sages)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	sagesMap := map[string]string{
		"India":     "Gandhi",
		"America":   "MLK",
		"Christian": "Jesus",
	}

	err := tpl.Execute(os.Stdout, sagesMap)
	if err != nil {
		log.Fatalln(err)
	}
}
