package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

type Response map[string]interface{}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl)
	t.Execute(w, p)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		p := &Page{Title: "Index"}
		renderTemplate(w, "views/index.html", p)
		// go queue(r.URL.Query()["tl"][0], r.URL.Query()["q"][0])
		// fmt.Fprint(w, Response{"success": true, "message": "Queued"})
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
