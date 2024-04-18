package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// a http server is created using "http.FileServer", which serves file form directory "./static"
	fileServer := http.FileServer(http.Dir("./static"))
	// the root path "/" is handled by fileServer
	http.Handle("/", fileServer)
	// two additional routes are setup using handleFunc
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting the server at 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
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
	fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parseForm() error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Adress: %s\n", address)
}
