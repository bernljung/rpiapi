package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages
	// send chan []byte
	send chan []message
}

func (c *connection) reader() {
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		var u utterance
		err = json.Unmarshal(msg, &u)
		if m, ok := u.validate(); ok {
			message := []message{{c: c, text: msg}}
			h.broadcast <- message
		} else {
			log.Printf("%v not valid: %v", string(msg), m)
		}
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		if err := c.ws.WriteMessage(websocket.TextMessage, message[0].text); err != nil {
			break
		}
	}
	c.ws.Close()
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func wsHandler(rw http.ResponseWriter, req *http.Request) {
	ws, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		return
	}
	c := &connection{send: make(chan []message), ws: ws}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}

func speakHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		rw.Header().Set("Content-Type", "application/json")
		u := utterance{
			Lang: req.FormValue("lang"),
			Text: req.FormValue("text"),
		}

		if m, ok := u.validate(); ok {
			log.Println("Text: ", u.Text)
			r := response{Success: ok, Message: "Here you go.", Text: u.Text, Lang: u.Lang}
			message := []message{{text: []byte(fmt.Sprintf("%v", r))}}
			h.broadcast <- message
			fmt.Fprint(rw, r)
		} else {
			fmt.Fprint(rw, response{Success: ok, Message: m})
		}
	} else {
		http.NotFound(rw, req)
	}
}
