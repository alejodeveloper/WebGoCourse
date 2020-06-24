package main

import (
	"html/template"
	"log"
	"net/http"
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
	myTemplate.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func init() {
	filePrefix, _ := filepath.Abs("./WebGo/HttpServer/formServerHttp/")
	myTemplate = template.Must(template.ParseFiles(filePrefix + "/index.gohtml"))
}

func main() {
	myServer := myServer{name: "Local"}
	http.ListenAndServe(":8080", myServer)
}
