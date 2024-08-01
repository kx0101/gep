package notifications

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func generateNotifications(now time.Time) []struct {
	Hour, Minute int
	Notification Notification
} {
	cleaningDescription := generateCleaningDescription(now.Weekday())
	flagHour, flagMinute = getFlagTimeForDate(now)

	dynamicNotifications := []struct {
		Hour, Minute int
		Notification Notification
	}{
		{
			cleaningHour, cleaningMinute, Notification{
				"Καθαριότητες Ταξιαρχίας",
				cleaningDescription,
				"audio/kathariotites-taksiarxias.mp3",
			},
		},
		{
			flagHour, flagMinute, Notification{
				"Σημαία",
				"Υποστολή Σημαίας, χωρίς jockey, κατεβάζουμε και λάβαρα",
				"audio/ypostolh-shmaias.mp3",
			},
		},
	}

	allNotifications := append(staticNotifications, dynamicNotifications...)

	return allNotifications
}

func playMP3(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {})))

	select {}
}

func generateCleaningDescription(weekday time.Weekday) string {
	cleaningInfo := cleaningSchedule[weekday]

	times := strings.Split(cleaningInfo.Duration, "-")
	startTime := strings.TrimSpace(times[0])
	hourMinute := strings.Split(startTime, ":")

	if len(hourMinute) == 2 {
		cleaningHour = parseTimePart(hourMinute[0])
		cleaningMinute = parseTimePart(hourMinute[1])
	}

	description := fmt.Sprintf("Καθαρίζουμε τα εξής: %s\nΔιάρκεια: %s", cleaningInfo.Places, cleaningInfo.Duration)
	return description
}

func getFlagTimeForDate(date time.Time) (int, int) {
	isDST := isDaylightSavingTime(date)

	flagSchedule := []struct {
		Start, End time.Time
		Time       FlagTime
	}{
		{time.Date(0, 1, 1, 0, 0, 0, 0, time.Local), time.Date(0, 1, 15, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{16, 50})},
		{time.Date(0, 1, 16, 0, 0, 0, 0, time.Local), time.Date(0, 1, 31, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{17, 10})},
		{time.Date(0, 2, 1, 0, 0, 0, 0, time.Local), time.Date(0, 2, 15, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{17, 25})},
		{time.Date(0, 2, 16, 0, 0, 0, 0, time.Local), time.Date(0, 2, 29, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{17, 40})},

		{time.Date(0, 3, 1, 0, 0, 0, 0, time.Local), time.Date(0, 3, 16, 0, 0, 0, 0, time.Local), adjustFlagTime(FlagTime{17, 55})},
		{time.Date(0, 3, 16, 0, 0, 0, 0, time.Local), time.Date(0, 3, lastSundayOfMonth(3, date.Year()).Day(), 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{18, 10})},
		{time.Date(0, 3, lastSundayOfMonth(3, date.Year()).Day()+1, 0, 0, 0, 0, time.Local), time.Date(0, 3, 31, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{19, 10})},

		{time.Date(0, 4, 1, 0, 0, 0, 0, time.Local), time.Date(0, 4, 15, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{19, 25})},
		{time.Date(0, 4, 16, 0, 0, 0, 0, time.Local), time.Date(0, 4, 30, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{19, 40})},
		{time.Date(0, 5, 1, 0, 0, 0, 0, time.Local), time.Date(0, 5, 15, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{19, 50})},
		{time.Date(0, 8, 1, 0, 0, 0, 0, time.Local), time.Date(0, 8, 15, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{20, 5})},
		{time.Date(0, 6, 1, 0, 0, 0, 0, time.Local), time.Date(0, 6, 15, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{20, 15})},
		{time.Date(0, 6, 16, 0, 0, 0, 0, time.Local), time.Date(0, 6, 30, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{20, 20})},
		{time.Date(0, 7, 1, 0, 0, 0, 0, time.Local), time.Date(0, 7, 15, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{20, 20})},
		{time.Date(0, 7, 16, 0, 0, 0, 0, time.Local), time.Date(0, 7, 31, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{20, 10})},
		{time.Date(0, 8, 1, 0, 0, 0, 0, time.Local), time.Date(0, 8, 15, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{20, 0})},
		{time.Date(0, 8, 16, 0, 0, 0, 0, time.Local), time.Date(0, 8, 31, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{19, 40})},
		{time.Date(0, 9, 1, 0, 0, 0, 0, time.Local), time.Date(0, 9, 15, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{19, 15})},
		{time.Date(0, 9, 16, 0, 0, 0, 0, time.Local), time.Date(0, 9, 30, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{18, 50})},

		{time.Date(0, 10, 1, 0, 0, 0, 0, time.Local), time.Date(0, 10, 16, 0, 0, 0, 0, time.Local), adjustFlagTime(FlagTime{18, 30})},
		{time.Date(0, 10, 16, 0, 0, 0, 0, time.Local), time.Date(0, 10, lastSundayOfMonth(10, date.Year()).Day(), 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{18, 05})},
		{time.Date(0, 10, lastSundayOfMonth(10, date.Year()).Day()+1, 0, 0, 0, 0, time.Local), time.Date(0, 10, 31, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{17, 05})},

		{time.Date(0, 11, 1, 0, 0, 0, 0, time.Local), time.Date(0, 11, 15, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{16, 50})},
		{time.Date(0, 11, 16, 0, 0, 0, 0, time.Local), time.Date(0, 11, 30, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{16, 40})},
		{time.Date(0, 12, 1, 0, 0, 0, 0, time.Local), time.Date(0, 12, 15, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{16, 35})},
		{time.Date(0, 12, 16, 0, 0, 0, 0, time.Local), time.Date(0, 12, 31, 23, 59, 59, 0, time.Local), adjustFlagTime(FlagTime{16, 40})},
	}

	for _, entry := range flagSchedule {
		if date.After(entry.Start) && date.Before(entry.End) {
			adjustedHour := entry.Time.Hour
			adjustedMinute := entry.Time.Minute

			if isDST && adjustedHour >= 18 {
				adjustedHour += 1
			}

			return adjustedHour, adjustedMinute
		}
	}

	return 0, 0
}

func adjustFlagTime(flagTime FlagTime) FlagTime {
	newMinute := flagTime.Minute - notifyBeforeMinutes
	newHour := flagTime.Hour

	if newMinute < 0 {
		newHour -= 1
		newMinute += 60
	}

	if newHour < 0 {
		newHour += 24
	}

	return FlagTime{Hour: newHour, Minute: newMinute}
}

func isDaylightSavingTime(date time.Time) bool {
	lastSundayMarch := lastSundayOfMonth(3, date.Year())
	lastSundayOctober := lastSundayOfMonth(10, date.Year())

	return date.After(lastSundayMarch) && date.Before(lastSundayOctober.Add(24*time.Hour))
}

func lastSundayOfMonth(month, year int) time.Time {
	t := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	t = t.Add(-24 * time.Hour)

	for t.Weekday() != time.Sunday {
		t = t.Add(-24 * time.Hour)
	}

	return t
}

func parseTimePart(part string) int {
	value, err := strconv.Atoi(part)
	if err != nil {
		fmt.Println("Error parsing time part:", err)
		return 0
	}

	return value
}
