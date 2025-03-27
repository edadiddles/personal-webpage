package server

import (
    "net/http"
    "fmt"
)

func Run() {

    htmlPages()
    jsFiles()
    cssFiles()

    fmt.Println("Listening on port 8080")
    http.ListenAndServe(":8080", nil)
}


func htmlPages() {
    http.HandleFunc("/home", homeHtmlHandler)
}

func jsFiles() {
    http.HandleFunc("/js/home.js", jsHomeFileHandler)
}

func cssFiles() {
    http.HandleFunc("/css/home.css", cssHomeFileHandler)
}

func homeHtmlHandler(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, "ui/html/home.html")
}

func jsHomeFileHandler(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, "ui/js/home.js")
}

func cssHomeFileHandler(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, "ui/css/home.css")
}

