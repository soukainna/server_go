// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// )

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got / request\n")
// 	io.WriteString(w, "This is my website!\n")
// }
// func getHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got /hello request\n")
// 	io.WriteString(w, "Hello, HTTP!\n")
// }

// func main() {

// 	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 	// 	fmt.Fprintf(w, "Hello!")
// 	// })

// 	http.HandleFunc("/", getRoot)
// 	http.HandleFunc("/hello", getHello)

// 	fmt.Printf("Starting server at port 8080\n")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
// 	// attention: If you do not call ParseForm method, the following data can not be obtained form
// 	fmt.Println(r.Form) // print information on server side.
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello soukaina!") // write data to response
// }

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {

		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password:", r.Form["password"])
		fmt.Fprintf(w, "Hello %s!", r.FormValue("username"))
	}
}

// func printName(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, welcome to %s!", r.URL.Path[4:])
// }

func main() {
	//http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)

	http.Handle(("/"), http.FileServer(http.Dir("./static")))
	//http.HandleFunc("/", printName)

	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
