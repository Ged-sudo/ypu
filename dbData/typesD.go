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

type ManhwaData struct {
	Id            uint16
	Name          string
	LinkImgManhwa []string
	Capture       string
}

type ManhwaTData struct {
	Id              uint16
	Name            string
	LincksImgManhwa string
	Capture         string
}
