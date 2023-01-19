package utils

import (
	"christmasfetch/colors"
	"christmasfetch/config"
	"math/rand"
)

var options []colors.Color

// Generates a light pattern with a random color
func GenerateLights(length int, config config.Config) string {
	options = append(options,
		colors.Blue,
		colors.Yellow,
		colors.Magenta,
		colors.Red,
		colors.Cyan,
	)

	var colorsList []colors.Color
	colorsList = append(colorsList, colors.Cyan, colors.Blue, colors.Magenta, colors.Red, colors.Yellow)
	


	var lights string
	switch config.Lights {
	case "individual":
		for i := 1; i <= length; i++ {
			chosenColor := colorsList[rand.Intn(len(colorsList))]
			lights += chosenColor.Color + "*" + colors.Reset.Color + "-"
		}
	default:
		found := false
		for option := range options {
			if config.Lights == options[option].Name {
				chosenColor := options[option].Color
				for i := 1; i <= length; i++ {
					lights += chosenColor + "*" + colors.Reset.Color + "-"
				}
				found = true
				break
			}
		}
		if found == false {
			chosenColor := colorsList[rand.Intn(len(colorsList))]
			for i := 1; i <= length; i++ {
				lights += chosenColor.Color + "*" + colors.Reset.Color + "-"
			}
			lights = lights + chosenColor.Color + "*" + colors.Reset.Color
		}
	}

	return lights
}
