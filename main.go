package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func apidocHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println(r.URL.Path)
		p := page{
			Title:    "API",
			Template: "public/templates/apidoc.html",
		}
		p.renderTemplate(w)
	} else {
		http.NotFound(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path[1:]) == 0 && r.Method == "GET" {
		log.Println(r.URL.Path)
		p := page{
			Title:    "What do you want me to say?",
			Template: "public/templates/index.html",
			Lang:     "sv",
		}
		p.renderTemplate(w)

	} else {
		http.NotFound(w, r)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path[1:])
	http.ServeFile(w, r, "public/"+r.URL.Path[1:])
}

func main() {
	flag.Parse()
	go h.run()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/static/", staticHandler)
	http.HandleFunc("/api/v1/docs", apidocHandler)
	http.HandleFunc("/api/v1/speak", speakHandler)
	http.HandleFunc("/ws", wsHandler)

	message := fmt.Sprintf("Starting server on %v", *addr)
	log.Println(message)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
