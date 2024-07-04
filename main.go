package main

import (
	"gep/tools"
	"sync"
)

func main() {
	// tools.CleanUp()

	var wg sync.WaitGroup
	wg.Add(4)

	go tools.CreateFireImage(&wg)
	go tools.CreateWeatherImageForToday(&wg)
	go tools.CreateWeatherImageForTomorrow(&wg)
	go tools.CreateWeatherImageForWeek(&wg)

	wg.Wait()

	tools.CreatePresentation()
}
