package weather

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

var (
	weatherTomorrow         = "https://freemeteo.gr/kairos/aleksandroupoli/oriaia-provlepsi/aurio/?gid=736928&language=greek&country=greece"
	weatherTwoDaysFromToday = "https://freemeteo.gr/kairos/aleksandroupoli/oriaia-provlepsi/day2/?gid=736928&language=greek&country=greece"
	weatherWeek             = "https://freemeteo.gr/kairos/aleksandroupoli/7-imeres/pinakas/?gid=736928&language=greek&country=greece"
	today                   = time.Now().Format("02-01-2006")
)

func CreateWeatherImageForTomorrow(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Fetching weather image for tomorrow...")

	tomorrow := time.Now().AddDate(0, 0, 1)
	filename := fmt.Sprintf("images/%s.png", tomorrow.Format("02-01-2006"))
	CreateWeatherImage(weatherTomorrow, filename)
}

func CreateWeatherImageForTwoDaysFromNow(wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Fetching weather image for two days from now...")

	twoDaysFromNow := time.Now().AddDate(0, 0, 2)
	filename := fmt.Sprintf("images/%s.png", twoDaysFromNow.Format("02-01-2006"))

	CreateWeatherImage(weatherTwoDaysFromToday, filename)
}

func CreateWeatherImageForWeek() {
	log.Printf("Fetching weather image for the week...")

	filename := fmt.Sprintf("images/ΕΒΔΟΜΑΔΙΑΙΟΣ - %s.png", today)
	CreateWeatherImage(weatherWeek, filename)
}

func CreateWeatherImage(url, filename string) {
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

	log.Printf("Screenshot saved as %s\n", filename)
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
