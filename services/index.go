package main

import (
	"log"
	"net/http"
	"text/template"
)

type val int

func (m val) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.html", req.Form)
}

var tpl *template.Template

func init() {
	//ParseFiles creates a new Template and parses the template definitions from the named file
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var d val
	http.ListenAndServe(":8080", d)
}
