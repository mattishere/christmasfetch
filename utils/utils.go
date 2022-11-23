package utils

import (
	"christmasfetch/colors"
	"math/rand"
)

// Generates a light pattern with a random color
func GenerateLights(length int) string {
	colorsList := []string{}
	colorsList = append(colorsList, colors.Cyan, colors.Blue, colors.Magenta, colors.Red, colors.Yellow)

	chosenColor := colorsList[rand.Intn(len(colorsList))]

	lights := ""
	for i := 1; i <= length; i++ {
		 lights = lights + chosenColor + "*" + colors.Reset + "-"
	}
	lights = lights + chosenColor + "*" + colors.Reset

	return lights
}
