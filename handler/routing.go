package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routing() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/videos", mainVideos)
	rtr.HandleFunc("/upload", uploadFile).Methods("POST")
	rtr.HandleFunc("/upload_video", uploadVideoF)
	rtr.HandleFunc("/posts", posts)
	rtr.HandleFunc("/vposts", vposts)
	rtr.HandleFunc("/upload_succes", succesUpload)
	rtr.HandleFunc("/video_view/{id:[0-9]+}", videoView)
	rtr.HandleFunc("/", koren)

	http.Handle("/", rtr)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../assets/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static/"))))
	http.ListenAndServe(":8080", nil)
}
