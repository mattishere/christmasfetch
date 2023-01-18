package utils

import "christmasfetch/colors"

type LightsOption struct {
	name string
	color string
}

var (
	Blue = LightsOption{"blue", colors.Blue}
	Yellow = LightsOption{"yellow", colors.Yellow}
	Green = LightsOption{"green", colors.Green}
	Magenta = LightsOption{"magenta", colors.Magenta}
	Red = LightsOption{"red", colors.Red}
	Cyan = LightsOption{"cyan", colors.Cyan}
)
