package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func main () {
	http.HandleFunc("/", python)
	http.ListenAndServe(":8080", nil)
}

func python (res http.ResponseWriter, req *http.Request) {
	filePrefix, _ := filepath.Abs("./WebGo/HttpServer/requestHttpServer/")
	f, err := os.Open(filePrefix + "/python.png")
	defer f.Close()
	if err != nil {
		http.Error(res, "File not found", 404)
		return
	}
	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "File not found", 404)
		return
	}
	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
}