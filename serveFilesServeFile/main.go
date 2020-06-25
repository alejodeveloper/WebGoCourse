package main

import (
	"net/http"
	"path/filepath"
)

func main () {
	http.HandleFunc("/", python)
	http.ListenAndServe(":8080", nil)
}

func python (res http.ResponseWriter, req *http.Request) {
	filePrefix, _ := filepath.Abs("./WebGo/HttpServer/requestHttpServer/")
	http.ServeFile(res, req, filePrefix + "python.png")
}