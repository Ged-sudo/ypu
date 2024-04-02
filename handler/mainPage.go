package handlers

import (
	"html/template"
	"main/db"
	dbdata "main/dbData"
	"net/http"
)

func mainVideos(w http.ResponseWriter, r *http.Request) {
	// templates/footer.html
	temp, err := template.ParseFiles("../templates/main_video.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}

	data = []dbdata.VideoData{}
	data = db.SpeshVideo()

	temp.ExecuteTemplate(w, "main_video", data)

}
