package main

import (
	"io"
	"net/http"
)

type myServer struct {

}

func (m myServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		io.WriteString(res, "Home")
	case "/dog":
		io.WriteString(res, "Dog")
	case "/cat":
		io.WriteString(res, "Cat")
	default:
		io.WriteString(res, "404")

	}
}

func main()  {
	var server myServer
	http.ListenAndServe(":8080", server)
}