package command

import (
	"math/rand"
	"time"
)


func Run() {
	// This makes every output random
	rand.Seed(time.Now().UnixNano())
	
	// Get the current date and the Christmas date
	currentDate := time.Now()
	christmasDate := GetChristmasDate(currentDate)

	// Get the days until Christmas
	days := DaysUntilChristmas(currentDate, christmasDate)

	// Check if it's Christmas (this adds an extra message when formatting if true)
	isItChristmas := false
	if days == 0 {
		isItChristmas = true
	}

	// Pass the info into the Info struct and use the Format() method to apply a theme and output the final result
	info := Info{christmasDate.Year(), christmasDate.Weekday().String(), currentDate.Format("Jan 02, 2006"), days, GiftIdea(), isItChristmas }
	Format(info)
}
