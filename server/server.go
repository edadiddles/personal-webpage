package server

import (
	"fmt"
	"net/http"
)

var port = ":8080"

func Run() {

	htmlPages()
	jsFiles()
	cssFiles()

	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(port, nil)
}

func htmlPages() {
	http.HandleFunc("/home", htmlHomeHandler)
	http.HandleFunc("/about", htmlAboutHandler)
	http.HandleFunc("/experience", htmlExperienceHandler)
	http.HandleFunc("/projects", htmlProjectsHandler)
	http.HandleFunc("/hobbies", htmlHobbiesHandler)
}

func jsFiles() {
	http.HandleFunc("/js/home.js", jsHomeHandler)
	http.HandleFunc("/js/about.js", jsAboutHandler)
	http.HandleFunc("/js/experience.js", jsExperienceHandler)
	http.HandleFunc("/js/projects.js", jsProjectsHandler)
	http.HandleFunc("/js/hobbies.js", jsHobbiesHandler)
}

func cssFiles() {
	http.HandleFunc("/css/home.css", cssHomeHandler)
	http.HandleFunc("/css/about.css", cssAboutHandler)
	http.HandleFunc("/css/experience.css", cssExperienceHandler)
	http.HandleFunc("/css/projects.css", cssProjectsHandler)
	http.HandleFunc("/css/hobbies.css", cssHobbiesHandler)
}

func htmlHomeHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/html/home.html")
}

func htmlAboutHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/html/about.html")
}

func htmlExperienceHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/html/experience.html")
}

func htmlProjectsHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/html/projects.html")
}

func htmlHobbiesHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/html/hobbies.html")
}


func jsHomeHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/js/home.js")
}

func jsAboutHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/js/about.js")
}

func jsExperienceHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/js/experience.js")
}

func jsProjectsHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/js/projects.js")
}

func jsHobbiesHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/js/hobbies.js")
}


func cssHomeHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/css/home.css")
}
func cssAboutHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/css/about.css")
}
func cssExperienceHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/css/experience.css")
}
func cssProjectsHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/css/projects.css")
}
func cssHobbiesHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "ui/css/hobbies.css")
}

