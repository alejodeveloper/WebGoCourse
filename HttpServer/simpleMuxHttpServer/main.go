package main

import (
	"io"
	"net/http"
)

type catPath struct {

}

type dogPath struct {

}

type homePath struct {

}

func (m dogPath) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog Dog")
}

func (m catPath) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Cat Cat")
}

func (m homePath) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Home")
}

func main()  {
	var dog dogPath
	var cat catPath
	var home homePath
	server := http.NewServeMux()
	server.Handle("/dog/", dog)
	server.Handle("/cat", cat)
	server.Handle("/", home)
	http.ListenAndServe(":8080", server)
}