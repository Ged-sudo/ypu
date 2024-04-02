package handlers

import (
	"fmt"
	"html/template"
	"main/db"
	dbdata "main/dbData"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../templates/index.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func video_view(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	t, err := template.ParseFiles("../templates/video_view.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}
	a := id["id"]
	fmt.Println("---------id---------")
	fmt.Println(a)
	fmt.Println("---------id---------")
	videoInfo := db.CurrentVideo(a)

	t.ExecuteTemplate(w, "video_view", videoInfo)
}

func posts(w http.ResponseWriter, r *http.Request) {

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
	// templates/posts.html
	// assets
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
	temp, err := template.ParseFiles("../templates/succes.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}
	temp.ExecuteTemplate(w, "succ", nil)
}

func koren(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../templates/koren.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}
	temp.ExecuteTemplate(w, "koren", nil)
}
