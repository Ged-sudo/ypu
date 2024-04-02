package db

import (
	"database/sql"
	"fmt"
	dbdata "main/dbData"
)

func CurrentVideo(id string) dbdata.VideoData {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/VideoHosting", user, pass, host, port))
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM Video WHERE id = %s;", id))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(res)

	var dataVideo = dbdata.VideoData{}
	err = res.Scan(&dataVideo.Id, &dataVideo.VideoHref, &dataVideo.ImgVideoHref, &dataVideo.NameVideo, &dataVideo.AuthorVideoName, &dataVideo.RangeIntresting)

	if err != nil {
		panic(err.Error())
	}

	return dataVideo
}

func SpeshVideo() []dbdata.VideoData {
	var videoList []dbdata.VideoData

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/VideoHosting", user, pass, host, port))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM Video")
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
