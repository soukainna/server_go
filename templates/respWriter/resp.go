package main

import (
	"fmt"
	"net/http"
)

//var tpl *template.Template

type res int

func (resp res) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `<h1>Hello</h1>`)
}

func main() {
	var w res
	http.ListenAndServe(":8080", w)
}
