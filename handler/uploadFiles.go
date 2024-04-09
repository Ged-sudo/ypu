package handlers

import (
	"fmt"
	"io/ioutil"
	"main/db"
	"main/parser"
	"net/http"
	"strconv"

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

func uploadManhwa(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	nameManhwa := r.FormValue("name_manhwa")
	linck := r.FormValue("linck")
	captureS := r.FormValue("capture_S")
	captureE := r.FormValue("capture_E")
	fmt.Println(linck)
	// "https://hmanga.org/manga/sextudy-group/глава-" + c + "/?style=list"
	// "https://hmanga.org/manga/sextudy-group/глава-"
	cs, err := strconv.Atoi(captureS)
	if err != nil {
		panic(err.Error())
	}

	ce, err := strconv.Atoi(captureE)
	if err != nil {
		panic(err.Error())
	}
	//https://hmanga.org/wp-content/uploads/WP-manga/data/manga_62e586377a5ce/88a2a885be48b1c09f44dac9637e7153/sexstudy_class_vol01_ch001_p002.jpg
	//https://hmanga.org/wp-content/uploads/WP-manga/data/manga_62e586377a5ce/ce0adc0dff76caf7e4896cd91c18eb96/sexstudy_class_vol01_ch003_p001.jpg
	//https://hmanga.org/wp-content/uploads/WP-manga/data/manga_62e586377a5ce/ec281880d63a7de9364f498bd73f8542/sexstudy_class_vol01_ch010_p001.jpg
	//https://hmanga.org/wp-content/uploads/WP-manga/data/manga_62e586377a5ce/88a2a885be48b1c09f44dac9637e7153/sexstudy_class_vol01_
	//ch001_p002.jpg
	if nameManhwa == "" || captureS == "" || captureE == "" || linck == "" {
		fmt.Fprintf(w, "Wrong data was input")
	} else {
		fmt.Println(cs, ce)
		for s := cs; s < ce; s++ {
			newS := strconv.Itoa(s)
			for i := 1; i < 8; i++ {
				newI := strconv.Itoa(i)
				pops := fmt.Sprintf(linck+"%s/p/%s", s, i)
				fmt.Println(pops)
				path, err := parser.ParseManhwa(pops, strconv.Itoa(s), strconv.Itoa(i))
				if err != nil {
					fmt.Println(s, i)
					break
				}
				db.AddManhwaToDB(nameManhwa, path, newS, newI)
			}
			// linck = linck + newS + "/?style=list"
			// fmt.Println(linck)
			// path := parser.ParseManhwa(linck, s)
		}
		http.Redirect(w, r, "/upload_succes", http.StatusSeeOther)
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
