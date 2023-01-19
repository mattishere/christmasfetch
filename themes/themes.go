package themes

import (
	"christmasfetch/colors"
	"christmasfetch/config"
	"christmasfetch/data"
	"christmasfetch/utils"
	"fmt"
	"math/rand"
	"strings"
)

var (
	themes = []Theme{}	
)

func Format(data data.ChristmasData, config config.Config) {
	lights := utils.GenerateLights(7, config)

	gift := Properties{primary: colors.Red, secondary: colors.Green, text: colors.White, lights: lights, data: data}
	tree := Properties{primary: colors.Green, secondary: colors.Red, text: colors.Yellow, lights: lights, data: data}
	candycane := Properties{primary: colors.Red, secondary: colors.White, text: colors.White, lights: lights, data: data}

	themes = append(themes, 
		Theme{name: "gift", art: Gift(gift)},
		Theme{name: "tree", art: Tree(tree)},
		Theme{name: "candycane", art: Candycane(candycane)},
	)

	switch config.Theme {
	case "random":
		fmt.Println(themes[rand.Intn(len(themes))].art)
	default:
		found := false
		for theme := range themes {
			if config.Theme == themes[theme].name {
				fmt.Println(themes[theme].art)
				found = true
				break
			}
		}
		if found == false {
			fmt.Println(Gift(gift))
		}
	}
	fmt.Println("\n\n")

	if data.IsChristmasDay == true {
		fmt.Println(colors.Red.Color + "		Mery christmas :)\n			- MattHere\n")
	}
}

type Properties struct {
	primary, secondary, text colors.Color
	lights string
	data data.ChristmasData
}

type Theme struct {
	name string
	art string
}

func Gift(props Properties) string {
	art :=
		`
	  ${PRIM}_
	  \\/_     ${SEC}Christmas${PRIM}@${SEC}${YEAR}
	${SEC}oooo${PRIM}|${SEC}oooo  ${LIGHTS}
	${SEC}osss${PRIM}|${SEC}ssso  ${PRIM}Is on: ${TEXT}${DAY}
	${PRIM}----|----  Is in: ${TEXT}${UNTIL} days
	${SEC}osss${PRIM}|${SEC}ssso  ${PRIM}Today is: ${TEXT}${DATE} 
	${SEC}oooo${PRIM}|${SEC}oooo  ${PRIM}Gift idea: ${TEXT}${GIFT}`

	final := parseArt(art, props)
	return final
}



func Tree(props Properties) string {
	art :=
		`
		 ${TEXT}*
		${PRIM}/~\
	       /~~~\	    ${SEC}Christmas${PRIM}@${SEC}${YEAR}
	      ${PRIM}/~~${SEC}o${PRIM}~~\	    ${LIGHTS}
	     ${PRIM}/~~~~~~~\	    ${PRIM}Is on: ${TEXT}${DAY}
	    ${PRIM}/~${SEC}o${PRIM}~~~~~${SEC}o${PRIM}~\	    ${PRIM}Is in: ${TEXT}${UNTIL} days
           ${PRIM}/~~~~~~~~~~~\    ${PRIM}Today is: ${TEXT}${DATE}
          ${PRIM}/~~~~${SEC}o${PRIM}~~~~~${SEC}o${PRIM}~~\   ${PRIM}Gift idea: ${TEXT}${GIFT}
         ${PRIM}/~${SEC}o${PRIM}~~~~~~~~~~~${SEC}o${PRIM}~\
	 ~~~~~~~~~~~~~~~~~		
		 `

	final := parseArt(art, props)
	return final
}

func Candycane(props Properties) string {
	art :=
	`
	${SEC} _${PRIM}_${SEC}_${PRIM}_	${PRIM}Christmas${SEC}@${PRIM}${YEAR}
	${SEC}/${PRIM}~${SEC}_${PRIM}_${SEC}~${PRIM}|	${LIGHTS}
	${PRIM}|${SEC}|  ${PRIM}|${SEC}|	${SEC}Is ${PRIM}on: ${SEC}${DAY}
	${SEC}|${PRIM}|	${PRIM}Is ${SEC}in: ${PRIM}${UNTIL} ${SEC}days
	${PRIM}|${SEC}|	${SEC}Today ${PRIM}is: ${SEC}${DATE}
	${SEC}|${PRIM}|	${PRIM}Gift ${SEC}idea: ${PRIM}${GIFT}
	`

	final := parseArt(art, props)
	return final
}



type Placeholder struct {
	placeholder string
	replacement string
}

func parseArt(art string, props Properties) string {
	// Setup all of the placeholders & their replacements
	var placeholders []Placeholder
	placeholders = append(placeholders, Placeholder{"${PRIM}", props.primary.Color})
	placeholders = append(placeholders, Placeholder{"${SEC}", props.secondary.Color})
	placeholders = append(placeholders, Placeholder{"${TEXT}", props.text.Color})
	
	placeholders = append(placeholders, Placeholder{"${LIGHTS}", props.lights})

	placeholders = append(placeholders, Placeholder{"${YEAR}", fmt.Sprint(props.data.Year)})
	placeholders = append(placeholders, Placeholder{"${DAY}", props.data.ChristmasDay})
	placeholders = append(placeholders, Placeholder{"${UNTIL}", fmt.Sprint(props.data.DaysUntil)})
	placeholders = append(placeholders, Placeholder{"${DATE}", props.data.CurrentDate})
	placeholders = append(placeholders, Placeholder{"${GIFT}", props.data.GiftIdea})

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
