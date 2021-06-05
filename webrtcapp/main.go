package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.New("").Delims("[[", "]]").ParseFiles("views/camera.html", "views/rtc.html", "views/peerjs.html", "views/call.html"))

func main() {
	http.HandleFunc("/camera", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "camera.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/rtc", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "rtc.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/peerjs", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "peerjs.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/call", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "call.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
