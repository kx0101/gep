package notifications

import "time"

type Notification struct {
	Title       string
	Description string
	Audio       string
}

type FlagTime struct {
	Hour   int
	Minute int
}

type CleaningDay struct {
	Places   string
	Duration string
}

const (
	dailyCleaningHours   = "16:00 - 18:00"
	weekendCleaningHours = "09:30 - 11:30"
	notifyBeforeMinutes  = 10
)

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
	{
		5, 45, Notification{
			Title:       "Εβδομαδιαίος Καιρός",
			Description: "Τοποθετείς τον εβδομαδιαίο καιρό που έχει δημιουργηθεί από το βραδινό script (gep/images) στο common/GEP/ΧΑΡΙΖΟΠΟΥΛΟΣ",
			Audio:       "audio/evdomadiaios-kairos.mp3",
		},
	},
	{
		12, 0, Notification{
			Title:       "Σημερινές Υπηρεσίες",
			Description: "Υπηρεσίες ΑΥΜΔ Μονάδων, Αξιωματικό πύλης/Μεσημβρίας, οδηγό επιφυλακής και Ν/Τ",
			Audio:       "audio/shmerines-yphresies.mp3",
		},
	},
	{
		15, 0, Notification{
			Title:       "Ετοίμασε τις υπηρεσίες για τον Συντονιστή",
			Description: "Άλλαξε το έγγραφο με το δελτίο υπηρεσιών της ταξιαρχίας",
			Audio:       "audio/yphresies-suntonisth.mp3",
		},
	},
	{
		18, 0, Notification{
			Title:       "Έλεγχος server room",
			Description: "Κατεβαίνεις κάτω στο server room και checkareis αν είναι όλα ΟΚ",
			Audio:       "audio/elegxos-server-room.mp3",
		},
	},
	{
		18, 10, Notification{
			Title:       "ΣΥΝΟΠΛΗ",
			Description: "Μπαίνεις στο ΠΥΡΣΕΙΑ του ΚΕΠΙΧ και κατεβάζεις το ΣΥΝΟΠΛΗ από την ΧΙΙ Μ/Κ. Έπειτα το περνάς common/GEP/ΧΑΡΙΖΟΠΟΥΛΟΣ",
			Audio:       "audio/sinopli.mp3",
		},
	},
	{
		19, 0, Notification{
			Title:       "Πότισμα",
			Description: "Γύρω από την ταξιαρχία και μπροστά από το κιόσκι",
			Audio:       "audio/potisma.mp3",
		},
	},
	{
		19, 30, Notification{
			Title:       "Επεισόδια ΕΕΑΣ, Μεσημβρίας και Άβαντα",
			Description: "Παίρνεις τηλέφωνο για τα επεισόδια",
			Audio:       "audio/epeisodia.mp3",
		},
	},
	{
		20, 15, Notification{
			Title:       "Σύνολα Εξοδούχων",
			Description: "Παίρνεις τηλέφωνο την Κεντρική Πύλη καθώς επίσης και 21 ΕΜΑ και 138 για να πάρεις τα σύνολα των εξοδούχων",
			Audio:       "audio/sunola-eksodouxwn.mp3",
		},
	},
	{
		21, 0, Notification{
			Title:       "Έλεγχος server room",
			Description: "Κατεβαίνεις κάτω στο server room και checkareis αν είναι όλα ΟΚ",
			Audio:       "audio/elegxos-server-room.mp3",
		},
	},
	{
		21, 10, Notification{
			Title:       "Αναφορά Συντονιστή",
			Description: "Ετοίμασε την αναφορά του συντονιστή",
			Audio:       "audio/anafora-suntonisth.mp3",
		},
	},
	{
		22, 0, Notification{
			Title:       "Αλλαγή powerpoint",
			Description: "Πας στο gep φάκελο στην επιφάνεια εργασίας του PC INTERNET και αφού ρυθμίσεις το loop στο powerpoint, το περνάς στο μαυρο usb",
			Audio:       "audio/allagh-powerpoint.mp3",
		},
	},
	{
		22, 45, Notification{
			Title:       "-1 ΚΣ",
			Description: "Είσαι ήδη -1 μέρα",
			Audio:       "audio/meiwn-ena.mp3",
		},
	},
}
