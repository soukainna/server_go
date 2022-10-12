// // package main

// // import (
// // 	"fmt"
// // 	"io"
// // 	"log"
// // 	"net/http"
// // )

// // func getRoot(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Printf("got / request\n")
// // 	io.WriteString(w, "This is my website!\n")
// // }
// // func getHello(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Printf("got /hello request\n")
// // 	io.WriteString(w, "Hello, HTTP!\n")
// // }

// // func main() {

// // 	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// // 	// 	fmt.Fprintf(w, "Hello!")
// // 	// })

// // 	http.HandleFunc("/", getRoot)
// // 	http.HandleFunc("/hello", getHello)

// // 	fmt.Printf("Starting server at port 8080\n")
// // 	if err := http.ListenAndServe(":8080", nil); err != nil {
// // 		log.Fatal(err)
// // 	}
// // }

// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"os"

// 	"git.com/soukainna/app"
// )

// /**************************************************************
// **************************************************************/

// // func sayhelloName(w http.ResponseWriter, r *http.Request) {
// // 	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
// // 	// attention: If you do not call ParseForm method, the following data can not be obtained form
// // 	fmt.Println(r.Form) // print information on server side.
// // 	fmt.Println("path", r.URL.Path)
// // 	fmt.Println("scheme", r.URL.Scheme)
// // 	fmt.Println(r.Form["url_long"])
// // 	for k, v := range r.Form {
// // 		fmt.Println("key:", k)
// // 		fmt.Println("val:", strings.Join(v, ""))
// // 	}
// // 	fmt.Fprintf(w, "Hello soukaina!") // write data to response
// // }

// func login(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("method:", r.Method) //get request method
// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("login.gtpl")
// 		t.Execute(w, nil)
// 	} else {

// 		r.ParseForm()
// 		// logic part of log in
// 		fmt.Println("username:", r.FormValue("username"))
// 		fmt.Println("password:", r.Form["password"])
// 		fmt.Fprintf(w, "Hello %s!", r.FormValue("username"))
// 	}
// }

// // func printName(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "Hi there, welcome to %s!", r.URL.Path[4:])
// // }

// /**************************************************************
// **************************************************************/
// type Page struct {
// 	Title string
// 	Body  []byte
// }

// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	return os.WriteFile(filename, p.Body, 0600)
// }

// func loadPage(title string) (*Page, error) {
// 	filename := title + ".txt"
// 	body, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	p, _ := loadPage(title)
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

// /**************************************************************
// **************************************************************/

// func main() {

// 	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
// 	// p1.save()
// 	// p2, _ := loadPage("TestPage")
// 	// fmt.Println(string(p2.Body))

// 	http.HandleFunc("/view/", viewHandler)

// 	fmt.Println("my name is", app.Nom)
// 	//fmt.Println(pack.hello)
// 	//http.HandleFunc("/", sayhelloName) // setting router rule
// 	http.HandleFunc("/login", login)

// 	http.Handle(("/"), http.FileServer(http.Dir("./static")))
// 	//http.HandleFunc("/", printName)

// 	err := http.ListenAndServe(":9090", nil) // setting listening port
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
