package handlers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

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

	if handlerImg.Filename[len(handlerImg.Filename)-3:] == "png" {
		tempFile, err := ioutil.TempFile("../assets/images", "upload-*.png")
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
		fmt.Println("Image .. done")
	}

	if handlerImg.Filename[len(handlerImg.Filename)-3:] == "jpg" {
		tempFile, err := ioutil.TempFile("../assets/images", "upload-*.jpg")
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
		fmt.Println("Image .. done")
	}

	if handlerImg.Filename[len(handlerImg.Filename)-4:] == "jpeg" {
		tempFile, err := ioutil.TempFile("../assets/images", "upload-*.jpeg")
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
		fmt.Println("Image .. done")
	}

	if handlerVideo.Filename[len(handlerVideo.Filename)-3:] == "m4v" {
		tempFile, err := ioutil.TempFile("../assets/videos", "upload-*.m4v")
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
		fmt.Println("Video .. done")
	}

	if handlerVideo.Filename[len(handlerVideo.Filename)-3:] == "mp4" {
		tempFile, err := ioutil.TempFile("../assets/videos", "upload-*.mp4")
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
		fmt.Println("Video .. done")
	}

	if nameVideo == "" || authorVideo == "" || rangeIntresting == "" {
		fmt.Fprintf(w, "Wrong data was input")
	} else {
		//TODO: add in db file this code
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/VideoHosting")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO Video(video_href, img_video_href, name_video, author_video_name, range_intresting) VALUES('%s', '%s', '%s', '%s', '%s');", video_href, img_video_href, nameVideo, authorVideo, rangeIntresting))
		if err != nil {
			panic(err)
		}

		defer insert.Close()
		// TODO: end code
		http.Redirect(w, r, "/upload_succes", http.StatusSeeOther)
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
