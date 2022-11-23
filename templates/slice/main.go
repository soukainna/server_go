package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	names := []string{"soukaina", "lol", "bob"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", names)
	if err != nil {
		log.Fatalln(err)
	}
}
