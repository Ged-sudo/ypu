package parser

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gocolly/colly"
)

func ParseManhwa(pageFromManhwa string, captureCount string, page string) (string, error) {
	path, err := scrap(pageFromManhwa, captureCount, page)
	return path, err
}

func scrap(scrapUrl string, capture string, page string) (string, error) {
	var linck string
	c := colly.NewCollector(colly.AllowedDomains("https://hmanga.org", "hmanga.org"))
	c.OnHTML("div.page-break", func(h *colly.HTMLElement) {
		linck = h.ChildAttr("img", "src")
		fmt.Println(linck) // добавить тут массив, потом при гете сделаать сортировку по порядку для корректности отображения...
	})
	c.Visit(scrapUrl)

	path, err := downloadFile("group", linck, capture, page)
	if err != nil {
		log.Fatal(err)
	}
	return path, err
}

// func ScrapImg(manhwaName string, URL string, capture int, page int) (string, error) {
// 	response, err := http.Get(URL)
// 	if err != nil {
// 		return "", err
// 	}

// 	defer response.Body.Close()

// 	if response.StatusCode != 200 {
// 		log.Fatal(err.Error())
// 	}
// 	//Create a empty file
// 	file, err := os.Create(fmt.Sprintf("../assets/manhwa/%s", fmt.Sprintf("%s_page_c%d_p%d.jpeg", manhwaName, capture, page)))
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	//Write the bytes to the fiel
// 	_, err = io.Copy(file, response.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	path := file.Name()

// 	return path, nil
// }

func downloadFile(manhwaName string, URL string, capture string, page string) (string, error) {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatal(err.Error())
	}
	//Create a empty file
	file, err := os.Create(fmt.Sprintf("../assets/manhwa/%s", fmt.Sprintf("%s_page_c%s_p%s.jpeg", manhwaName, capture, page)))
	if err != nil {
		return "", err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	path := file.Name()

	return path, nil
}
