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
	p1 := person{Name: "lol", Age: 22}
	p2 := person{Name: "bob", Age: 32}
	p3 := person{Name: "wow", Age: 15}
	p4 := person{Name: "sali", Age: 31}
	p5 := person{Name: "mami", Age: 68}

	persons := []person{p1, p2, p3, p4, p5}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", persons)
	if err != nil {
		log.Fatalln(err)
	}
}
