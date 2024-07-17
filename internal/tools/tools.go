package tools

import (
	"gep/internal/tools/presentation/fire"
	"gep/internal/tools/presentation/presentation"
	"gep/internal/tools/presentation/weather"
	"sync"
)

func CreatePresentation() {
	var wg sync.WaitGroup
	wg.Add(3)

	go fire.CreateFireImage(&wg)
	go weather.CreateWeatherImageForTomorrow(&wg)
	go weather.CreateWeatherImageForTwoDaysFromNow(&wg)

	wg.Wait()

	presentation.CreatePresentation()
}

func CreateWeeklyWeather() {
	weather.CreateWeatherImageForWeek()
}
