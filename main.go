package main

import (
	"flag"
	"fmt"
	"gep/internal/tools"
	"log"
)

func main() {
	fmt.Printf(`
    GEP Internal Tools:


    1 - Generate fire, weather images and create the presentation.
    2 - Generate weekly weather.
    3 - Notifications service for GEP guidance.


    ** Default option: 1**
    `)
	fmt.Println()

	option := flag.Int("tool", 1, "Select an option for which tool to use")
	flag.Parse()

	switch *option {
	case 1:
		tools.CreatePresentation()
	case 2:
		tools.CreateWeeklyWeather()
	case 3:
		tools.StartNotificationsService()
	default:
		log.Println("Invalid choice")
		return
	}
}
