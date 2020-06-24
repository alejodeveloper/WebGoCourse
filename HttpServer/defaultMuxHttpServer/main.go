package main

import (
	"io"
	"net/http"
)

func dog (res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog Dog")
}

func cat (res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Cat Cat")
}

func home (res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Home")
}

func main()  {

	// With the HandleFunc of http you dont need to pass a Handler but a handler who manage the path
	// handler different than a Handler
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat", cat)
	http.HandleFunc("/", home)
	// When nil is pass to ListenAndServe func it will implement defaultMuxServer
	http.ListenAndServe(":8080", nil)
}