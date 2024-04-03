package main

import (
	"fmt"
	"main/db"
	dbdata "main/dbData"
	"testing"
)

func TestCurrentVideo(t *testing.T) {
	id := "1"
	expected := dbdata.VideoData{
		Id:              1,
		VideoHref:       "../assets/videos/upload-2026346843.mp4",
		ImgVideoHref:    "../assets/images/upload-2286776434.jpg",
		NameVideo:       "1",
		AuthorVideoName: "11",
		RangeIntresting: "111",
	}

	actual, err := db.CurrentVideo(id)

	if err != nil {
		t.Errorf("Empty file")
	}

	if expected != actual {
		t.Errorf("Result was incorrect, got: %s; Want: %s", actual.NameVideo, expected.NameVideo)
	}
}

func TestSpeshVideo(t *testing.T) {
	test_v := []dbdata.VideoData{}
	expected := fmt.Sprintf("%T", test_v)

	actual := db.SpeshVideo()
	actual_type := fmt.Sprintf("%T", actual)

	if expected != actual_type {
		t.Errorf("Result wasincorrect, got: %s; Want: %s", actual_type, expected)
	}
}
