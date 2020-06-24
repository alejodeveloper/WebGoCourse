// Problem https://github.com/GoesToEleven/golang-web-dev/tree/master/022_hands-on/01/05_hands-on

package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
)

var myTemplate *template.Template


func dog (res http.ResponseWriter, req *http.Request) {
	io.WriteString(res,"Dog page")
}

func me (res http.ResponseWriter, req *http.Request) {
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
	myTemplate.ExecuteTemplate(res, "index.gohtml", data)
}

func init() {
	filePrefix, _ := filepath.Abs("./WebGo/exercises/exercisesOne/ninjaLevelOne3/")
	myTemplate = template.Must(template.ParseFiles(filePrefix + "/index.gohtml"))
}

func main () {
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}