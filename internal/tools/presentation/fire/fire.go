package fire

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	fireImageUrl = ""
)

func CreateFireImage(wg *sync.WaitGroup) {
	defer wg.Done()

	fireImageUrl, err := GetFireImageUrl()
	if err != nil {
		log.Fatalf("Error getting fire image url: %v", err)
	}

	DownloadAndSaveFireImage(fireImageUrl)
}

func GetFireImageUrl() (string, error) {
	log.Printf("Fetching url of civil protection...")
	response, err := http.Get("https://civilprotection.gov.gr/arxeio-imerision-xartwn")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Error: received non-200 status code %d while fetching image", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalf("Error loading HTML: %v", err)
	}

    // CHANGE
	tomorrow := time.Now().AddDate(0, 0, 0).Format("02/01/2006")
	fireImageSrc := ""
	doc.Find("a.maps_tile").Each(func(i int, s *goquery.Selection) {
		if s.AttrOr("data-sub-html", "") == tomorrow {
			img := s.Find("img#myimage")
			fireImageSrc, _ = img.Attr("src")
		}
	})

	if fireImageSrc == "" {
		fmt.Println("Image not found")
		return "", fmt.Errorf("image not found")
	}

	return "https://civilprotection.gov.gr" + fireImageSrc, nil
}

func DownloadAndSaveFireImage(fireImageUrl string) {
	log.Printf("Fetching fire image...")
	imgResponse, err := http.Get(fireImageUrl)
	if err != nil {
		log.Fatalf("Error fetching image: %v", err)
	}

	defer imgResponse.Body.Close()

	if imgResponse.StatusCode != 200 {
		log.Fatalf("Error: received non-200 status code %d while fetching image", imgResponse.StatusCode)
	}

	filename := fmt.Sprintf("images/%s-fire.jpg", time.Now().AddDate(0, 0, 1).Format("02-01-2006"))

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	defer file.Close()

	_, err = io.Copy(file, imgResponse.Body)
	if err != nil {
		log.Fatalf("Error saving image: %v", err)
	}

	log.Printf("Fire image saved as %s", filename)
}
