package command

import (
	"christmasfetch/config"
	"christmasfetch/data"
	"christmasfetch/themes"
	"flag"
	"math/rand"
	"time"
)


func Run() {
	genConfig := flag.Bool("gen-config", false, "Generate the config (if not present).")

	flag.Parse()

	if *genConfig == true {
		config.GenerateConfig()
		return
	}

	config := config.GetConfig()

	rand.Seed(time.Now().UnixNano())
	
	// Get the current date and the Christmas date
	currentDate := time.Now()
	christmasDate := data.GetChristmasDate(currentDate)

	// Get the days until Christmas
	days := data.DaysUntilChristmas(currentDate, christmasDate)

	// Check if it's Christmas (this adds an extra message when formatting if true)
	IsItChristmas := false
	if days == 0 {
		IsItChristmas = true
	}

	// Pass the info into the Info struct and use the Format() method to apply a theme and output the final result
	info := data.ChristmasData{christmasDate.Year(), christmasDate.Weekday().String(), currentDate.Format("Jan 02, 2006"), days, data.GiftIdea(), IsItChristmas}
	themes.Format(info, config)
}
