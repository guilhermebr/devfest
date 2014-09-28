package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	var IndexVars struct {
		Title string
	}

	IndexVars.Title = "Google DevFest Centro Oeste"
	err := templates.ExecuteTemplate(w, "index.html", IndexVars)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("views/*.html"))

	http.Handle("/img/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", IndexHandler)

	fmt.Println("Rodando em 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
