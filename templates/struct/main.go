package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	me := person{Name: "lol", Age: 22}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", me)
	if err != nil {
		log.Fatalln(err)
	}
}
