package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.New("").Delims("[[", "]]").ParseFiles("views/taking_still_photos_with_webrtc.html", "views/a_simple_rtcdatachannel_sample.html", "views/chat_for_peerjs.html", "views/call_for_peerjs.html"))

func main() {

	// WebRTCで静止画を撮る。以下のチュートリアルを参考に実装。
	//https://developer.mozilla.org/en-US/docs/Web/API/WebRTC_API/Taking_still_photos
	http.HandleFunc("/taking_still_photos_with_WebRTC", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "taking_still_photos_with_webrtc.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//  WebRTC API の RTCDataChannel を使ったサンプル。以下のチュートリアルを参考に実装。
	// https://developer.mozilla.org/en-US/docs/Web/API/WebRTC_API/Simple_RTCDataChannel_sample
	http.HandleFunc("/a_simple_rtcdatachannel_sample", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "a_simple_rtcdatachannel_sample.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// peerjs の DataConnection を使用してチャットをする
	http.HandleFunc("/chat_for_peerjs", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "chat_for_peerjs.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// peerjs の MediaConnection を使用してテレビ電話をする
	http.HandleFunc("/call_for_peerjs", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "call_for_peerjs.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
