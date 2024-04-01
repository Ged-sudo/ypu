package handlers

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
