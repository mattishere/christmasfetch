package command

import (
	"christmasfetch/colors"
	"christmasfetch/utils"
	"fmt"
	"strings"
)

func Format(info Info) {
	lights := utils.GenerateLights(7)
	gift := Gift(Theme{colors.Red, colors.Green, colors.White, lights, info})

	fmt.Println(gift + "\n\n")
}

type Theme struct {
	primary string
	secondary string
	text string
	lights string
	data Info
}

func Gift(theme Theme) string {
	art :=
		`
	  ${PRIM}_
	  \\/_    ${SEC}Christmas${PRIM}@${SEC}${YEAR}
	${SEC}oooo${PRIM}|${SEC}oooo  ${LIGHTS}
	${SEC}osss${PRIM}|${SEC}ssso  ${PRIM}Is on: ${TEXT}${DAY}
	${PRIM}----|----  Is in: ${TEXT}${UNTIL} days
	${SEC}osss${PRIM}|${SEC}ssso  ${PRIM}Today is: ${TEXT}${DATE} 
	${SEC}oooo${PRIM}|${SEC}oooo  ${PRIM}Gift idea: ${TEXT}${GIFT}`

	final := parseArt(art, theme)
	return final
}

type Placeholder struct {
	placeholder string
	replacement string
}


func parseArt(art string, theme Theme) string {
	// Setup all of the placeholders & their replacements
	var placeholders []Placeholder
	placeholders = append(placeholders, Placeholder{"${PRIM}", theme.primary})
	placeholders = append(placeholders, Placeholder{"${SEC}", theme.secondary})
	placeholders = append(placeholders, Placeholder{"${TEXT}", theme.text})
	
	placeholders = append(placeholders, Placeholder{"${LIGHTS}", theme.lights})

	placeholders = append(placeholders, Placeholder{"${YEAR}", fmt.Sprint(theme.data.year)})
	placeholders = append(placeholders, Placeholder{"${DAY}", theme.data.christmasDay})
	placeholders = append(placeholders, Placeholder{"${UNTIL}", fmt.Sprint(theme.data.daysUntil)})
	placeholders = append(placeholders, Placeholder{"${DATE}", theme.data.currentDate})
	placeholders = append(placeholders, Placeholder{"${GIFT}", theme.data.giftIdea})

	parsedArt := art
	for i := range placeholders {
		curr := placeholders[i]
		parsedArt = strings.ReplaceAll(parsedArt, curr.placeholder, curr.replacement)	
	}
	return parsedArt
}
