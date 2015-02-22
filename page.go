package main

import (
	"html/template"
	"net/http"
)

type page struct {
	Title, Flash, Lang, Text, Template string
	Success                            bool
}

func (p *page) renderTemplate(w http.ResponseWriter) {
	t, _ := template.ParseFiles(p.Template)
	t.Execute(w, p)
}
