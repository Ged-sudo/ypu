package handlers

import (
	"fmt"
	"html/template"
	"log"
	"main/db"
	dbdata "main/dbData"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func uploadVideoF(w http.ResponseWriter, r *http.Request) {
	log.Println("User on upload page (uploadVideoF)")
	t, err := template.ParseFiles("../templates/index.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func videoView(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	log.Printf("User on current video page (videoView) id = %s", id["id"])

	t, err := template.ParseFiles("../templates/video_view.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}
	a := id["id"]
	videoInfo, err := db.CurrentVideo(a)
	if err != nil {
		fmt.Printf(err.Error())
	}

	t.ExecuteTemplate(w, "video_view", videoInfo)
}

func posts(w http.ResponseWriter, r *http.Request) {
	log.Println("Check all imgs in path (posts)")
	t, err := template.ParseFiles("../templates/posts.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}

	dir, err := os.Open("../assets/images")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()

	// Получаем список файлов и папок
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	var linkImg = []dbdata.IgmHrefs{}
	for _, file := range files {
		var name dbdata.IgmHrefs
		name.HrefImg = file.Name()
		linkImg = append(linkImg, name)
	}

	t.ExecuteTemplate(w, "posts", linkImg)
}

func vposts(w http.ResponseWriter, r *http.Request) {
	log.Println("Check videos in path(vposts)")

	t, err := template.ParseFiles("../templates/posts.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}

	dir, err := os.Open("../assets/videos") // /Users/evgenii/Desktop/GolangPAth/GoTestEducation/ypu/assets/videos
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	var linkVid = []dbdata.VidHrefs{}
	for _, file := range files {
		var name dbdata.VidHrefs
		name.HrefVid = file.Name()
		linkVid = append(linkVid, name)
	}

	t.ExecuteTemplate(w, "vposts", linkVid)
}

var data = []dbdata.VideoData{}

func succesUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("User on succes upload page (succesUpload)")
	temp, err := template.ParseFiles("../templates/succes.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}
	temp.ExecuteTemplate(w, "succ", nil)
}

func koren(w http.ResponseWriter, r *http.Request) {
	log.Println("User '/' (koren)")
	temp, err := template.ParseFiles("../templates/koren.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}
	temp.ExecuteTemplate(w, "koren", nil)
}
