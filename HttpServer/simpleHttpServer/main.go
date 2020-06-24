package main

import (
	"fmt"
	"net/http"
)

type myStruct struct {
	some string
}

func (m myStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "My first HTTP Server on GO")
}

func main() {
	var server myStruct
	http.ListenAndServe(":8080", server)
}
