package dbdata

type IgmHrefs struct {
	HrefImg string
	Author  string
}

type VidHrefs struct {
	HrefVid string
	Author  string
}

type VideoData struct {
	Id              uint16
	VideoHref       string
	ImgVideoHref    string
	NameVideo       string
	AuthorVideoName string
	RangeIntresting string
}
