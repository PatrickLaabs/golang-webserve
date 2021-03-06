package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", hom)
	http.HandleFunc("/about", abo)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))
	http.Handle("/filestwo/", http.StripPrefix("/filestwo/", http.FileServer(http.Dir("./filestwo"))))
	http.ListenAndServe(":8200", nil)
}

func hom(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "default.gohtml", nil)
}

func abo(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}
