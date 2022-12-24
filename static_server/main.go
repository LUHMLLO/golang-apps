package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succesful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileserver := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	parsedUrl, _ := url.Parse("http://localhost:3080/")
	fmt.Println("\n Starting server at")
	fmt.Println("\n", parsedUrl)

	if err := http.ListenAndServe(":3080", nil); err != nil {
		log.Fatal(err)
	}
}
