package themes

import (
	"christmasfetch/colors"
	"christmasfetch/data"
	"christmasfetch/files"
	"christmasfetch/utils"
	"fmt"
	"math/rand"
	"strings"
)

func Format(data data.ChristmasData, config files.Config) {
	lights := utils.GenerateLights(7, config)

	themes := files.GetThemes()

	gift := files.Theme{Name: "gift", Art: `    ${RED}_
     \\/_    ${GREEN}Christmas${RED}@${GREEN}${YEAR}
  ${GREEN}oooo${RED}|${GREEN}oooo  ${LIGHTS}
  ${GREEN}osss${RED}|${GREEN}ssso  ${RED}Is on: ${WHITE}${DAY}
  ${RED}----|----  Is in: ${WHITE}${UNTIL} days
  ${GREEN}osss${RED}|${GREEN}ssso  ${RED}Today is: ${WHITE}${DATE} 
  ${GREEN}oooo${RED}|${GREEN}oooo  ${RED}Gift idea: ${WHITE}${GIFT}`}

	themes = append(themes)

	theme := data.Theme
	if theme == "" {
		theme = config.Theme
	}

    fmt.Println() // Adds a new line at the top.
	switch theme {
	case "random":
		themes = append(themes, gift)
		selectedTheme := themes[rand.Intn(len(themes))].Art
		final := parseArt(selectedTheme, data, lights)
		fmt.Println(final)
	default:
		found := false
		// Since Gift is the only internal theme, we want to hard-code it in.
		if len(themes) == 0 {
			fmt.Println(parseArt(gift.Art, data, lights))
		} else if theme == "gift" {
			fmt.Println(parseArt(gift.Art, data, lights))
		} else {
			for i := range themes {
				if theme == themes[i].Name {
					fmt.Println(parseArt(themes[i].Art, data, lights))
					found = true
					break
				}
			}
			if found == false {
				fmt.Println(parseArt(gift.Art, data, lights))
			}
		}

	}
	fmt.Println() // Adds a new line at the bottom.

	if data.IsChristmasDay == true {
		fmt.Println(colors.Red.Color + "		Mery christmas :)\n			- MattHere\n" + colors.Reset.Color)
	}
}

type Properties struct {
	primary, secondary, text colors.Color
	lights                   string
	data                     data.ChristmasData
}

type Placeholder struct {
	placeholder string
	replacement string
}

func parseArt(art string, data data.ChristmasData, lights string) string {
	// Setup all of the placeholders & their replacements
	var placeholders []Placeholder
	placeholders = append(placeholders, Placeholder{"${LIGHTS}", lights})

	placeholders = append(placeholders, Placeholder{"${YEAR}", fmt.Sprint(data.Year)})
	placeholders = append(placeholders, Placeholder{"${DAY}", data.ChristmasDay})
	placeholders = append(placeholders, Placeholder{"${UNTIL}", fmt.Sprint(data.DaysUntil)})
	placeholders = append(placeholders, Placeholder{"${DATE}", data.CurrentDate})
	placeholders = append(placeholders, Placeholder{"${GIFT}", data.GiftIdea})

	for color := range colors.ColorsList {
		placeholders = append(placeholders, Placeholder{"${" + strings.ToUpper(colors.ColorsList[color].Name) + "}", colors.ColorsList[color].Color})
	}

	parsedArt := art
	for i := range placeholders {
		curr := placeholders[i]
		parsedArt = strings.ReplaceAll(parsedArt, curr.placeholder, curr.replacement)
	}
	return parsedArt
}
