package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var everyspeakerPort *int

type Page struct {
	Title, Flash, Lang, Speech string
	Success                    bool
}

type JSONResponse struct {
	Success bool
	Message string
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
	t, _ := template.ParseFiles("public/" + tmpl)
	t.Execute(w, p)
}

func validate(lang, speech string) (string, bool) {
	if speech != "" {
		if lang == "sv" || lang == "en" {
			return "Nemas problemas", true
		} else {
			return "Only 'sv' and 'en' are appropriate values for 'lang'.", false
		}
	} else {
		return "You need to add a value for 'speech'.", false
	}
}

func postToEverySpeaker(lang, speech string) (string, bool) {
	everyspeakerAddress := fmt.Sprintf("http://localhost:%v/post", *everyspeakerPort)
	if res, err := http.PostForm(everyspeakerAddress,
		url.Values{"tl": {lang}, "q": {speech}}); err == nil {
		defer res.Body.Close()
		if body, err := ioutil.ReadAll(res.Body); err == nil {
			var jsonRes JSONResponse
			json.Unmarshal(body, &jsonRes)
			if jsonRes.Success == true {
				return "A speaker will get to work immediately.", true
			} else {
				return "Pft, my speakers seems to have gone completely insane.", false
			}
		} else {
			return "Oops, all my speakers seems to have gone fishin'.", false
		}
	} else {
		return "Oops, all my speakers seems to have gone fishin'.", false
	}
}

func speakHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		lang := r.FormValue("lang")
		speech := r.FormValue("speech")

		if message, valid := validate(lang, speech); valid {
			message, success := postToEverySpeaker(lang, speech)
			fmt.Fprint(w, Response{"success": success, "message": message})
		} else {
			fmt.Fprint(w, Response{"success": false, "message": "Hmpf, I deserve better input than that. " + message})
		}

	} else {
		http.NotFound(w, r)
	}
}

func apidocHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{}
	if r.Method == "GET" {
		log.Println(r.URL.Path)
		p = &Page{Title: "API"}
		renderTemplate(w, "templates/apidoc.html", p)
	} else {
		http.NotFound(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{}
	if len(r.URL.Path[1:]) == 0 {
		switch r.Method {
		case "GET":
			log.Println(r.URL.Path)
			p = &Page{Title: "What do you want me to say?"}
			renderTemplate(w, "templates/index.html", p)

		case "POST":
			lang := r.FormValue("lang")
			speech := r.FormValue("speech")

			if message, valid := validate(lang, speech); valid {
				if message, error := postToEverySpeaker(lang, speech); error {
					p = &Page{Flash: message, Lang: lang, Success: true}
				} else {
					p = &Page{Flash: message, Lang: lang, Speech: speech, Success: false}
				}
			} else {
				p = &Page{Flash: message, Lang: lang, Speech: speech, Success: false}
			}

			p.Title = "What do you want me to say?"
			renderTemplate(w, "templates/index.html", p)

		default:
			http.NotFound(w, r)
		}
	} else {
		http.NotFound(w, r)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path[1:])
	http.ServeFile(w, r, "public/"+r.URL.Path[1:])
}

func main() {
	port := flag.Int("port", 8080, "listen port number")
	everyspeakerPort = flag.Int("everyspeakerPort", 8000, "everyspeaker port number")
	flag.Parse()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/static/", staticHandler)
	http.HandleFunc("/api/v1/docs", apidocHandler)
	http.HandleFunc("/api/v1/speak", speakHandler)

	message := fmt.Sprintf("Starting server on :%v", *port)
	log.Println(message)
	address := fmt.Sprintf(":%v", *port)
	log.Fatal(http.ListenAndServe(address, nil))
}
