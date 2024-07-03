package fire

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	fireImageUrl = ""
	today        = time.Now().Format("02/01/2006")
)

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

	fireImageSrc := ""
	doc.Find("a.maps_tile").Each(func(i int, s *goquery.Selection) {
		if s.AttrOr("data-sub-html", "") == today {
			img := s.Find("img#myimage")
			fireImageSrc, _ = img.Attr("src")
		}
	})

	if fireImageSrc == "" {
		fmt.Println("Image not found")
		return "", fmt.Errorf("image not found")
	}

	log.Printf("Image source created")
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

	today = time.Now().Format("02-01-2006")
	filename := fmt.Sprintf("images/%s-fire.jpg", today)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, imgResponse.Body)
	if err != nil {
		log.Fatalf("Error saving image: %v", err)
	}

	fmt.Printf("Fire image saved as %s", filename)
    fmt.Println()
}
