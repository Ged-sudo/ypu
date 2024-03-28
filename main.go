package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type igmHrefs struct {
	HrefImg string
	Author  string
}

type vidHrefs struct {
	HrefVid string
	Author  string
}

type videoData struct {
	Id              uint16
	VideoHref       string
	ImgVideoHref    string
	NameVideo       string
	AuthorVideoName string
	RangeIntresting string
}

var linkImg = []igmHrefs{}
var linkVid = []vidHrefs{}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	var img_video_href string
	var video_href string
	// r.ParseMultipartForm(100 << 20)
	nameVideo := r.FormValue("video_name")
	authorVideo := r.FormValue("author")
	rangeIntresting := r.FormValue("range_intresting")

	fileVideo, handlerVideo, err := r.FormFile("my_file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	fileImg, handlerImg, err := r.FormFile("image_video")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer fileImg.Close()
	defer fileVideo.Close()

	fmt.Printf("Uploaded File: %+v\n", handlerVideo.Filename)

	fmt.Printf("File Size: %+v\n", handlerVideo.Size)
	fmt.Printf("MIME Header: %+v\n", handlerVideo.Header)

	if handlerImg.Filename[len(handlerImg.Filename)-3:] == "png" {
		tempFile, err := ioutil.TempFile("assets/images", "upload-*.png")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		img_video_href = tempFile.Name()
		fileBytes, err := ioutil.ReadAll(fileImg)
		if err != nil {
			fmt.Println(err)
		}

		tempFile.Write(fileBytes)
	}

	if handlerVideo.Filename[len(handlerVideo.Filename)-3:] == "m4v" {
		tempFile, err := ioutil.TempFile("assets/videos", "upload-*.m4v")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(fileVideo)
		if err != nil {
			fmt.Println(err)
		}
		video_href = tempFile.Name()
		tempFile.Write(fileBytes)
	}

	if handlerVideo.Filename[len(handlerVideo.Filename)-3:] == "mp4" {
		tempFile, err := ioutil.TempFile("assets/videos", "upload-*.mp4")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(fileVideo)
		if err != nil {
			fmt.Println(err)
		}
		video_href = tempFile.Name()
		tempFile.Write(fileBytes)
	}

	if nameVideo == "" || authorVideo == "" || rangeIntresting == "" {
		fmt.Fprintf(w, "Wrong data was input")
	} else {

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:33060)/VideoHosting")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO Video(video_href, img_video_href, name_video, author_video_name, range_intresting) VALUES('%s', '%s', '%s', '%s', '%s');", video_href, img_video_href, nameVideo, authorVideo, rangeIntresting))
		if err != nil {
			panic(err)
		}

		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
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
	linkImg = []igmHrefs{}
	for _, file := range files {
		var name igmHrefs
		name.HrefImg = file.Name()
		linkImg = append(linkImg, name)
	}

	t.ExecuteTemplate(w, "posts", linkImg)
}

func vposts(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/vpost.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err.Error())
	}

	dir, err := os.Open("./assets/videos")
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
	linkVid = []vidHrefs{}
	for _, file := range files {
		var name vidHrefs
		name.HrefVid = file.Name()
		linkVid = append(linkVid, name)
	}

	t.ExecuteTemplate(w, "vposts", linkVid)
}

var data = []videoData{}

func mainVideos(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/main_video.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:33060)/VideoHosting")
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

func setupRoutes() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/videos", mainVideos)
	rtr.HandleFunc("/upload", uploadFile).Methods("POST")
	rtr.HandleFunc("/upload_video", mainPage)
	rtr.HandleFunc("/posts", posts)
	rtr.HandleFunc("/vposts", vposts)

	http.Handle("/", rtr)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Hello World")

	setupRoutes()
}
