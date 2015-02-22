package main

import (
	"bytes"
	"encoding/json"
)

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Text    string `json:"text"`
	Lang    string `json:"lang"`
}

func (r response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	s = string(b)
	return
}
