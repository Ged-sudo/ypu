package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../templates/index.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func posts(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/posts.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err.Error())
	}

	dir, err := os.Open("./assets/images")
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
	var linkImg = []igmHrefs{}
	for _, file := range files {
		var name igmHrefs
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

	dir, err := os.Open("/Users/evgenii/Desktop/GolangPAth/GoTestEducation/ypu/assets/videos") // /Users/evgenii/Desktop/GolangPAth/GoTestEducation/ypu/assets/videos
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
	var linkVid = []vidHrefs{}
	for _, file := range files {
		var name vidHrefs
		name.HrefVid = file.Name()
		linkVid = append(linkVid, name)
	}

	t.ExecuteTemplate(w, "vposts", linkVid)
}

var data = []videoData{}

func mainVideos(w http.ResponseWriter, r *http.Request) {
	// templates/footer.html
	temp, err := template.ParseFiles("../templates/main_video.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		panic(err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/VideoHosting")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM Video")
	if err != nil {
		panic(err.Error())
	}

	data = []videoData{}
	for res.Next() {
		var dataVideo videoData
		err = res.Scan(&dataVideo.Id, &dataVideo.VideoHref, &dataVideo.ImgVideoHref, &dataVideo.NameVideo, &dataVideo.AuthorVideoName, &dataVideo.RangeIntresting)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, dataVideo)
	}

	temp.ExecuteTemplate(w, "main_video", data)

}
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
