// Problem https://github.com/GoesToEleven/golang-web-dev/tree/master/022_hands-on/01/01_hands-on

package main

import (
	"io"
	"net/http"
)

func dog (res http.ResponseWriter, req *http.Request) {
	io.WriteString(res,"Dog page")
}

func me (res http.ResponseWriter, req *http.Request) {
	io.WriteString(res,"Alex page")
}

func main () {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}