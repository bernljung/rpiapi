package main

var VALID_LANGS = [...]string{
	"de-DE", "da-DK", "en-US", "en-GB", "es-ES",
	"fi-FI", "fr-FR", "nb-NO", "ru-RU", "sv-SE",
}

type utterance struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

func (u utterance) validate() (string, bool) {
	if u.Text != "" {
		for i := 0; i < len(VALID_LANGS); i++ {
			if u.Lang == VALID_LANGS[i] {
				return "Validated.", true
			}
		}
		return "You need to add an appropriate value for lang.", false
	} else {
		return "You need to add appropriate value for text.", false
	}
}
