package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type form int

func (f form) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "form.html", r.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("form.html"))
}

func main() {

	var fm form
	http.ListenAndServe(":8080", fm)

}
