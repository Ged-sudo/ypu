package db

import (
	"database/sql"
	"fmt"
	dbdata "main/dbData"
)

func CurrentVideo(id string) dbdata.VideoData {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, nameDB))
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `%s` WHERE `id` = '%s';", nameTable, id))
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

	return dataVideoOne[0]
}

func SpeshVideo() []dbdata.VideoData {
	var videoList []dbdata.VideoData

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, nameDB))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM %s", nameTable))
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

	insert, err := db.Query(fmt.Sprintf("INSERT INTO %s(video_href, img_video_href, name_video, author_video_name, range_intresting) VALUES('%s', '%s', '%s', '%s', '%s');", nameTable, video_href, img_video_href, nameVideo, authorVideo, rangeIntresting))
	if err != nil {
		panic(err)
	}

	defer insert.Close()
}
