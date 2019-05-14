package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

var quotes = []struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}{
	{"If you cannot fail, you cannot learn.", "Eric Ries"},
	{"If Tetris has taught me anything, it's that errors pile up and accomplishments disappear.", "Andrew Clay Shafer"},
	{"The most dangerous phrase in the language is \"we've always done it this way\".", "Rear Admiral Grace Hopper"},
	{"Perhaps instead of \"we must get better at estimates\" we should try \"let's become less dependent on fortune telling.\"", "Andy Hunt"},
	{"Our customers are loyal to us right until the second somebody offers them a better service.", "Jeff Bezos	"},
	{"The illiterate of the 21st century will not be those who cannot read and write, but those who cannot learn, unlearn, and relearn.", "Alvin Toffler"},
}

func main() {

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", QuoteHandler)

	fmt.Println(http.ListenAndServe(":8080", nil))
}

// QuoteHandler returns a random quote selected from the list
func QuoteHandler(w http.ResponseWriter, r *http.Request) {
	index := rand.Intn(len(quotes) - 1)

	js, err := json.Marshal(quotes[index])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
