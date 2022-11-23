package command

import (
	"christmasfetch/colors"
	"christmasfetch/utils"
	"fmt"
)

func Format(info Info) {
	// TODO: Implement support for multiple formats (randomly chosen)
	Present(info)
}




// TODO: Implement a cleaner and easier way to add formats

func Present(info Info) {
	art :=
	`
    %s_
     \\/_    %sChristmas%s@%s%s
  %soooo%s|%soooo  %s
  %sosss%s|%sssso  %sIs on: %s%s
  %s----|----  Is in: %s%s days
  %sosss%s|%sssso  %sToday is: %s%s 
  %soooo%s|%soooo  %sGift idea: %s%s

  %s%s`


    // Merry Christmas :)
	merryChristmasMessage := ""
	if(info.isChristmasDay == true) {
		merryChristmasMessage = colors.Red + "	Merry Christmas :)" + colors.Reset + "\n		- MattIsHere"
	}


	fmt.Printf(art,
			   colors.Red,
			   colors.Green, colors.Red, colors.Green, fmt.Sprint(info.year),
			   colors.Green, colors.Red, colors.Green, utils.GenerateLights(7),
			   colors.Green, colors.Red, colors.Green, colors.Red, colors.White, info.christmasDay,
			   colors.Red, colors.White, fmt.Sprint(info.daysUntil),
			   colors.Green, colors.Red, colors.Green, colors.Red, colors.White, info.currentDate,
			   colors.Green, colors.Red, colors.Green, colors.Red, colors.White, info.giftIdea,
			   fmt.Sprint(merryChristmasMessage), string("\n\n"))
}
