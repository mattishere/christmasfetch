package utils

import (
	"christmasfetch/colors"
	"christmasfetch/config"
	"math/rand"
)

var options []LightsOption

// Generates a light pattern with a random color
func GenerateLights(length int, config config.Config) string {
	options = append(options,
		Blue,
		Yellow,
		Magenta,
		Red,
		Cyan,
	)

	colorsList := []string{}
	colorsList = append(colorsList, colors.Cyan, colors.Blue, colors.Magenta, colors.Red, colors.Yellow)
	


	var lights string
	switch config.Lights {
	case "individual":
		for i := 1; i <= length; i++ {
			chosenColor := colorsList[rand.Intn(len(colorsList))]
			lights += chosenColor + "*" + colors.Reset + "-"
		}
	default:
		found := false
		for option := range options {
			if config.Lights == options[option].name {
				chosenColor := options[option].color
				for i := 1; i <= length; i++ {
					lights += chosenColor + "*" + colors.Reset + "-"
				}
				found = true
				break
			}
		}
		if found == false {
			chosenColor := colorsList[rand.Intn(len(colorsList))]
			for i := 1; i <= length; i++ {
				lights += chosenColor + "*" + colors.Reset + "-"
			}
			lights = lights + chosenColor + "*" + colors.Reset
		}
	}

	return lights
}
