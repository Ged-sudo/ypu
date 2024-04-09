package handlers

import (
	"fmt"
	"html/template"
	"log"
	"main/db"
	"net/http"

	"github.com/gorilla/mux"
)

func manhwa(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)
	name_manhwa := data["name-manhwa"]
	capture := data["capture"]

	log.Println(fmt.Printf("User on upload page (manhwa, %s, %s)", name_manhwa, capture))

	t, err := template.ParseFiles("../templates/manhwa.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}

	result := db.Manhwa(capture, name_manhwa)

	t.ExecuteTemplate(w, "manhwa", result)
}
