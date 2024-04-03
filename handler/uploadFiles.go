package handlers

import (
	"fmt"
	"io/ioutil"
	"main/db"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	nameVideo := r.FormValue("video_name")
	authorVideo := r.FormValue("author")
	rangeIntresting := r.FormValue("range_intresting")

	current_img_path := ImgCurrentType(r)
	video_href := VideoCurrentPath(r)

	if nameVideo == "" || authorVideo == "" || rangeIntresting == "" {
		fmt.Fprintf(w, "Wrong data was input")
	} else {
		db.AddVideoToDB(video_href,
			current_img_path,
			nameVideo,
			authorVideo,
			rangeIntresting,
		)
		http.Redirect(w, r, "/upload_succes", http.StatusSeeOther)
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func VideoCurrentPath(r *http.Request) string {
	var video_path string
	fileVideo, handlerVideo, err := r.FormFile("my_file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)

	}

	defer fileVideo.Close()

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
		video_path = tempFile.Name()
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
		video_path = tempFile.Name()
		tempFile.Write(fileBytes)
		fmt.Println("Video .. done")
	}

	return video_path
}

func ImgCurrentType(r *http.Request) string {

	var img_video_href string

	fileImg, handlerImg, err := r.FormFile("image_video")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		// return
	}

	defer fileImg.Close()

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

	return img_video_href
}
