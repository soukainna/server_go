package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello")
	http.HandleFunc("/home", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/refresh", Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
