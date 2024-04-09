package db

import (
	"database/sql"
	"fmt"
	dbdata "main/dbData"
)

func CurrentVideo(id string) (dbdata.VideoData, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, nameDB))
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `%s` WHERE `id` = '%s';", nameVideoTable, id))
	if err != nil {
		panic(err.Error())
	}

	var dataVideoOne = []dbdata.VideoData{}

	for res.Next() {
		var dataVideo dbdata.VideoData
		err = res.Scan(&dataVideo.Id, &dataVideo.VideoHref, &dataVideo.ImgVideoHref, &dataVideo.NameVideo, &dataVideo.AuthorVideoName, &dataVideo.RangeIntresting)
		if err != nil {
			panic(err.Error())
		}
		dataVideoOne = append(dataVideoOne, dataVideo)
	}

	if err != nil {
		panic(err.Error())
	}

	return dataVideoOne[0], nil
}

func SpeshVideo() []dbdata.VideoData {
	var videoList []dbdata.VideoData

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, nameDB))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM %s", nameVideoTable))
	if err != nil {
		panic(err.Error())
	}

	videoList = []dbdata.VideoData{}
	for res.Next() {
		var dataVideo dbdata.VideoData
		err = res.Scan(&dataVideo.Id, &dataVideo.VideoHref, &dataVideo.ImgVideoHref, &dataVideo.NameVideo, &dataVideo.AuthorVideoName, &dataVideo.RangeIntresting)
		if err != nil {
			panic(err.Error())
		}
		videoList = append(videoList, dataVideo)
	}

	return videoList
}

func AddVideoToDB(video_href string, img_video_href string, nameVideo string, authorVideo string, rangeIntresting string) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, nameDB))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO %s(video_href, img_video_href, name_video, author_video_name, range_intresting) VALUES('%s', '%s', '%s', '%s', '%s');", nameVideoTable, video_href, img_video_href, nameVideo, authorVideo, rangeIntresting))
	if err != nil {
		panic(err)
	}

	defer insert.Close()
}

func AddManhwaToDB(name string, linckImg string, capture string, page string) {

	nameM := name
	linckImgM := linckImg
	captureM := capture

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, nameDB))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO %s(Name, linckImg, capture, page) VALUES('%s', '%s', '%s', '%s')", nameManhwaTable, nameM, linckImgM, captureM, page))
	if err != nil {
		panic(err)
	}

	defer insert.Close()
}

//SELECT * FROM Mnhwa WHERE capture = `123`  AND Name = `asd`

func Manhwa(capture string, nameM string) dbdata.ManhwaData {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, nameDB))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE capture = '%s' AND Name = '%s'", nameManhwaTable, capture, nameM))
	if err != nil {
		panic(err.Error())
	}

	var arr []string
	for res.Next() {
		var manhwaDataM dbdata.ManhwaTData
		err = res.Scan(&manhwaDataM.Id, &manhwaDataM.Name, &manhwaDataM.LincksImgManhwa, &manhwaDataM.Capture)
		if err != nil {
			panic(err.Error())
		}
		arr = append(arr, manhwaDataM.LincksImgManhwa)

	}
	fmt.Println(arr)
	var result = dbdata.ManhwaData{
		Id:            0,
		Name:          nameM,
		LinkImgManhwa: arr,
		Capture:       capture,
	}

	return result
}
