package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

type Notification struct {
	Title       string
	Description string
}

type FlagTime struct {
	Hour   int
	Minute int
}

type CleaningDay struct {
	Places   string
	Duration string
}

var dailyCleaningHours = "16:00 - 18:00"
var weekendCleaningHours = "09:30 - 11:30"
var notifyBefore = 10
var cleaningHour, cleaningMinute int
var flagHour, flagMinute int

var cleaningSchedule = map[time.Weekday]CleaningDay{
	time.Monday: {
		Places:   "Γραφείο ΥΔΚΤΗ, 1ο(2 ΓΡΑΦΕΙΑ), 2ο(2 ΓΡΑΦΕΙΑ), 3ο(3 ΓΡΑΦΕΙΑ)",
		Duration: dailyCleaningHours,
	},
	time.Tuesday: {
		Places:   "Γραφείο ΕΠΧΗ, 4ο(4 ΓΡΑΦΕΙΑ), ΓΕΠ(2 ΓΡΑΦΕΙΑ), ΓΡΑΜΜΑΤΕΙΑ (1 ΓΡΑΦΕΙΟ)",
		Duration: dailyCleaningHours,
	},
	time.Wednesday: {
		Places:   "ΚΟΙΝΟΧΡΗΣΤΟ ΧΩΡΟΙ, ΚΕΠΙΧ",
		Duration: dailyCleaningHours,
	},
	time.Thursday: {
		Places:   "Γραφείο ΥΔΚΤΗ, 1ο(2 ΓΡΑΦΕΙΑ), 2ο(2 ΓΡΑΦΕΙΑ), 3ο(3 ΓΡΑΦΕΙΑ)",
		Duration: dailyCleaningHours,
	},
	time.Friday: {
		Places:   "Γραφείο ΕΠΧΗ, 4ο(4 ΓΡΑΦΕΙΑ), ΓΕΠ(2 ΓΡΑΦΕΙΑ), ΓΡΑΜΜΑΤΕΙΑ (1 ΓΡΑΦΕΙΟ)",
		Duration: dailyCleaningHours,
	},
	time.Saturday: {
		Places:   "ΧΩΡΟΙ ΣΤΡΑΤΗΓΕΙΟΥ",
		Duration: weekendCleaningHours,
	},
	time.Sunday: {
		Places:   "ΑΙΘΟΥΣΑ ΕΠΙΧΕΙΡΗΣΕΩΝ",
		Duration: weekendCleaningHours,
	},
}

var staticNotifications = []struct {
	Hour, Minute int
	Notification Notification
}{
	{5, 45, Notification{"Εβδομαδιαίος Καιρός", "Τοποθετείς τον εβδομαδιαίο καιρό που έχει δημιουργηθεί από το βραδινό script (gep/images) στο common/GEP/ΧΑΡΙΖΟΠΟΥΛΟΣ"}},
    {7, 22, Notification{"Έπαρση Σημαίας", "Σημαία, και το jockey στην τσέπη. Λάβαρα μετά την έπαρση"}},
	{7, 50, Notification{"Έπαρση Σημαίας", "Σημαία, και το jockey στην τσέπη. Λάβαρα μετά την έπαρση"}},
	{12, 0, Notification{"Σημερινές Υπηρεσίες", "Υπηρεσίες ΑΥΜΔ Μονάδων, Αξιωματικό πύλης/Μεσημβρίας, οδηγό επιφυλακής και Ν/Τ"}},
	{15, 0, Notification{"Ετοίμασε τις υπηρεσίες για τον Συντονιστή", "Άλλαξε το έγγραφο με το δελτίο υπηρεσιών της ταξιαρχίας"}},
	{18, 0, Notification{"Έλεγχος server room", "Κατεβαίνεις κάτω στο server room και checkareis αν είναι όλα ΟΚ"}},
	{18, 10, Notification{"ΣΥΝΟΠΛΗ", "Μπαίνεις στο ΠΥΡΣΕΙΑ του ΚΕΠΙΧ και κατεβάζεις το ΣΥΝΟΠΛΗ από την ΧΙΙ Μ/Κ. Έπειτα το περνάς common/GEP/ΧΑΡΙΖΟΠΟΥΛΟΣ"}},
	{19, 0, Notification{"Πότισμα", "Γύρω από την ταξιαρχία και μπροστά από το κιόσκι"}},
	{19, 30, Notification{"Επεισόδια ΕΕΑΣ, Μεσημβρίας και Άβαντα", "Παίρνεις τηλέφωνο για τα επεισόδια"}},
	{20, 15, Notification{"Σύνολα Εξοδούχων", "Παίρνεις τηλέφωνο την Κεντρική Πύλη καθώς επίσης και 21 ΕΜΑ και 138 για να πάρεις τα σύνολα των εξοδούχων"}},
	{21, 0, Notification{"Έλεγχος server room", "Κατεβαίνεις κάτω στο server room και checkareis αν είναι όλα ΟΚ"}},
	{21, 10, Notification{"Αναφορά Συντονιστή", "Ετοίμασε την αναφορά του συντονιστή"}},
	{22, 0, Notification{"Αλλαγή powerpoint", "Πας στο gep φάκελο στην επιφάνεια εργασίας του PC INTERNET και αφού ρυθμίσεις το loop στο powerpoint, το περνάς στο μαυρο usb"}},
	{22, 45, Notification{"-1 ΚΣ", "Είσαι ήδη -1 μέρα"}},
}

func adjustFlagTime(flagTime FlagTime) FlagTime {
	newMinute := flagTime.Minute - notifyBefore
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

func isNotificationTime(now time.Time) (*Notification, bool) {
	allNotifications := generateNotifications(now)

	for _, entry := range allNotifications {
		if now.Hour() == entry.Hour && now.Minute() == entry.Minute {
			return &entry.Notification, true
		}
	}

	return nil, false
}

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
		{cleaningHour, cleaningMinute, Notification{"Καθαριότητες Ταξιαρχίας", cleaningDescription}},
		{flagHour, flagMinute, Notification{"Σημαία", "Υποστολή Σημαίας, χωρίς jockey, κατεβάζουμε και λάβαρα"}},
	}

	allNotifications := append(staticNotifications, dynamicNotifications...)

	return allNotifications
}

func sendNotification(notification *Notification) {
	err := beeep.Alert(notification.Title, notification.Description, "")
	if err != nil {
		fmt.Println("Failed to send notification:", err)
	}
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

func main() {
	fmt.Println("Notification service started...")

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for now := range ticker.C {
		info, notify := isNotificationTime(now)

		if notify {
			sendNotification(info)
		}
	}
}
