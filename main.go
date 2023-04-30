package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Form struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(res, "Hello World")
}

func formHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseFrom() err: %v", err)
		return
	}
	fmt.Fprintf(res, "POST request successful\n")
	name := req.FormValue("name")
	address := req.FormValue("address")

	fmt.Fprintf(res, "Name = %s\n", name)
	fmt.Fprintf(res, "Address = %s\n", address)
}
func jsonHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/json" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "Method is not supported", http.StatusNotFound)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	data := &Form{
		Name:    "John",
		Address: "New York",
	}
	d, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	res.Write(d)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/json", jsonHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
