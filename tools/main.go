package tools

import (
	"gep/tools/fire"
	"gep/tools/presentation"
	"gep/tools/weather"
	"log"
	"sync"
)

func CreateFireImage(wg *sync.WaitGroup) {
	defer wg.Done()

	fireImageUrl, err := fire.GetFireImageUrl()
	if err != nil {
		log.Fatalf("Error getting fire image url: %v", err)
	}

	fire.DownloadAndSaveFireImage(fireImageUrl)
}

func CreateWeatherImageForToday(wg *sync.WaitGroup) {
	defer wg.Done()

	weather.CreateWeatherImageForToday()
}

func CreateWeatherImageForTomorrow(wg *sync.WaitGroup) {
	defer wg.Done()

	weather.CreateWeatherImageForTomorrow()
}

func CreateWeatherImageForWeek(wg *sync.WaitGroup) {
	defer wg.Done()

	weather.CreateWeatherImageForWeek()
}

func CreatePresentation() {
	presentation.CreatePresentation()
}
