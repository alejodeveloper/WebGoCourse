package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
)

var myTemplate *template.Template

type myServer struct {
	name string
}

func (m myServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method string
		URL *url.URL
		Submissions map[string][]string
		Header http.Header
		Host string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	myTemplate.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	filePrefix, _ := filepath.Abs("./WebGo/HttpServer/requestHttpServer/")
	myTemplate = template.Must(template.ParseFiles(filePrefix + "/index.gohtml"))
}

func main() {
	myServer := myServer{name: "Local"}
	http.ListenAndServe(":8080", myServer)
}
