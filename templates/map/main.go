package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	address := map[string]string{
		"soukaina": "Cannes",
		"lol":      "Nice",
		"bob":      "Paris",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", address)
	if err != nil {
		log.Fatalln(err)
	}
}
