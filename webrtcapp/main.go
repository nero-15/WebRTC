package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.New("").Delims("[[", "]]").ParseFiles("views/camera.html", "views/rtc.html", "views/peerjs.html", "views/call.html"))

func main() {

	// WebRTCで静止画を撮る。以下のチュートリアルを参考に実装。
	//https://developer.mozilla.org/en-US/docs/Web/API/WebRTC_API/Taking_still_photos
	http.HandleFunc("/take_photo_with_WebRTC", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "camera.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//  WebRTC API の RTCDataChannel を使ったサンプル。以下のチュートリアルを参考に実装。
	// https://developer.mozilla.org/en-US/docs/Web/API/WebRTC_API/Simple_RTCDataChannel_sample
	http.HandleFunc("/RTCDataChannel_sample", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "rtc.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/peerjs", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplae(w, "peerjs.html", nil)
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
