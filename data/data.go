package data

import (
	"math"
	"math/rand"
	"time"
)

type ChristmasData struct {
	Year			int
	ChristmasDay	string
	CurrentDate		string 
	DaysUntil		int 
	GiftIdea		string
	IsChristmasDay	bool
    Theme           string
}


func GetChristmasDate(currentTime time.Time) time.Time {
	year := currentTime.Year()

	today := time.Date(year, currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)
	date := time.Date(year, 12, 25, 0, 0, 0, 0, time.Local)

	if today.After(date) {
		date = time.Date(year + 1, 12, 25, 0, 0, 0, 0, time.Local)
	}

	return date
}

func DaysUntilChristmas(currentTime time.Time, christmasTime time.Time) int {
	// By subtracting the two dates, we can get the hours and divide them by 24 to get days until Christmas
	days := christmasTime.Sub(currentTime).Hours() / 24

	// This makes it into a valid integer by removing the decimals
	daysInt := math.Floor(days)

	if daysInt == -1 {
		daysInt = 0
	}

	return int(daysInt)
}

func GiftIdea() string {
	// I really tried keeping this list organised but I couldn't keep up...
	giftsList := []string{}
	giftsList = append(giftsList,
		"a laptop",
		"a computer",
		"a book",
		"a notebook",
		"sweets",
		"candy canes",
		"a guitar",
		"guitar picks",
		"a cookbook",
		"cooking pots",
		"cooking utensils",
		"a reading lamp",
		"a cool pen",
		"led strips",
		"a custom keyboard",
		"a plant",
		"gardening supplies",
		"a bonsai tree",
		"a personalized keychain",
		"a calendar",
		"headphones",
		"a smartwatch",
		"a bike",
		"a water bottle",
		"a pillow",
		"board games",
		"games",
		"speakersfalse",
		"a microphone",
		"smart home stuff",
		"a hoodie",
		"pants",
		"a hat",
		"a santa cap",
		"a wireless charger",
		"a kitten",
		"a puppy",
		"a guitar pedal",
		"a guitar amplifier",
		"a piano",
		"a plant pot",
		"sunglasses",
		"a deskmat",
		"a chessboard",
		"cards",
		"weights",
		"a great carpet",
		"diy stuff",
		"an awesome shirt",
		"a sweater",
		"a backpack",
		"an acquarium",
		"a raspberry pi",
		"discord nitro (don't pls)",
		"a vacation",
		"a hamster",
		"a lizard",
		"a fox",
		"an alpaca",
		"a llama",
		"a controller",
		"an online course",
		"a gaming console",
		"a software license",
		"a flashlight",
		"a multitool",
		"a blanket",
		"a new phone",
		"keyboard switches",
		"keycaps",
		"a capybara",
		// This is where I was really out of ideas
		"a robot",
		"a car",
		"land")


	// Return a randomly selected gift idea
	return giftsList[rand.Intn(len(giftsList))]
}
