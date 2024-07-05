package weather

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

var (
	weatherToday    = "https://freemeteo.gr/kairos/aleksandroupoli/oriaia-provlepsi/simera/?gid=736928&language=greek&country=greece"
	weatherTomorrow = "https://freemeteo.gr/kairos/aleksandroupoli/oriaia-provlepsi/aurio/?gid=736928&language=greek&country=greece"
	weatherWeek     = "https://freemeteo.gr/kairos/aleksandroupoli/7-imeres/pinakas/?gid=736928&language=greek&country=greece"
	today           = time.Now().Format("02-01-2006")
	filename_GEP    = "C:\\GEP Data Downloader\\Downloaded Files\\weather"
)

func CreateWeatherImageForToday() {
	log.Printf("Fetching weather image for today...")

	filename := fmt.Sprintf("images/%s.png", today)
	createWeatherImage(weatherToday, filename)
}

func CreateWeatherImageForTomorrow() {
	log.Printf("Fetching weather image for tomorrow...")

	tomorrow := time.Now().AddDate(0, 0, 1)
	filename := fmt.Sprintf("images/%s.png", tomorrow.Format("02-01-2006"))

	createWeatherImage(weatherTomorrow, filename)
}

func CreateWeatherImageForWeek() {
	log.Printf("Fetching weather image for the week...")

	filename := fmt.Sprintf("images/ΕΒΔΟΜΑΔΙΑΙΟΣ - %s.png", today)
	createWeatherImage(weatherWeek, filename)
}

func createWeatherImage(url, filename string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx, fullScreenshot(url, ".weather-now", &buf))
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(filename, buf, 0644); err != nil {
		log.Fatal(err)
	}

    if err := os.WriteFile(filename_GEP, buf, 0644); err != nil {
        log.Fatal(err)
    }

	fmt.Printf("Screenshot saved as %s\n", filename)
}

func fullScreenshot(url, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(sel, chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector(".fc-consent-root").remove()`, nil),
		chromedp.Evaluate(`document.querySelectorAll(".weather-now .graph-btn").forEach(e => e.remove())`, nil),
		chromedp.Evaluate(`document.querySelectorAll(".weather-now .graph").forEach(e => e.remove())`, nil),
		chromedp.Evaluate(`document.querySelector(".table-menu").remove()`, nil),
		chromedp.Sleep(time.Second * 2),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}
