// Problem https://github.com/GoesToEleven/golang-web-dev/tree/master/022_hands-on/01/03_hands-on

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
	filePrefix, _ := filepath.Abs("./WebGo/exercises/exercisesOneOne/ninjaLevelOne2/")
	myTemplate = template.Must(template.ParseFiles(filePrefix + "/index.gohtml"))
}

func main () {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}