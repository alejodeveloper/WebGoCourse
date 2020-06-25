package main

import (
	"net/http"
	"path/filepath"
)

func main () {
	http.Handle("/", http.FileServer(http.Dir("./serveFilesServeFiles/")))
	http.HandleFunc("/python/", python)
	http.ListenAndServe(":8080", nil)
}

func python (res http.ResponseWriter, req *http.Request) {
	filePrefix, _ := filepath.Abs("./serveFilesServeFiles/")
	http.ServeFile(res, req, filePrefix + "/python.png")
}