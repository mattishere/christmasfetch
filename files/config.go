package files

import (
	"christmasfetch/colors"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Theme	string `json:"name"`
	Lights	string `json:"lights"`
}

var defaultData = Config{
	Theme: "random",
	Lights: "individual",
}

func GetConfig() Config {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(configDir, "christmasfetch", "config.json")

	if _, err := os.Stat(path); err != nil {
		return defaultData
	}

	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}


	var config Config
	if err := json.Unmarshal(content, &config); err != nil {
		log.Fatal(err)
	}

	GetThemes()

	return config
}

func GenerateConfig() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(configDir, "christmasfetch")

	if _, err := os.Stat(path); err != nil {
		err := os.Mkdir(path, 0700)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(colors.Green.Color + "[+] Created the configuration directory at " + colors.White.Color + path + colors.Reset.Color)
		}
	}
	
	if _, err := os.Stat(filepath.Join(path, "themes")); err != nil {
		err = os.Mkdir(filepath.Join(path, "themes"), 0700)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(colors.Green.Color + "[+] Created the themes directory at " + colors.White.Color + filepath.Join(path, "themes") + colors.Reset.Color)
		}
	}

	tree := []byte(`     ${YELLOW}*     ${RED}Christmas${GREEN}@${RED}${YEAR}
    ${GREEN}/~\    ${LIGHTS}
   ${GREEN}/~~${MAGENTA}o${GREEN}\   Is on: ${YELLOW}${DAY}
  ${GREEN}/~${CYAN}o${GREEN}~~~\  Is in: ${YELLOW}${UNTIL} days
 ${GREEN}/~~~~~${RED}o${GREEN}~\ Today is: ${YELLOW}${DATE}
    ${YELLOW}|||    ${GREEN}Gift idea: ${YELLOW}${GIFT}`)
	makeFile(filepath.Join(path, "themes", "tree.theme"), tree)

	candycane := []byte(`  ${WHITE} _${RED}_${WHITE}_${RED}_	  ${RED}Christmas${WHITE}@${RED}${YEAR}
  ${WHITE}/${RED}~${WHITE}_${RED}_${WHITE}~${RED}|  ${LIGHTS}
  ${RED}|${WHITE}|  ${RED}|${WHITE}|  ${WHITE}Is ${RED}on: ${WHITE}${DAY}
  ${WHITE}|${RED}|	  ${RED}Is ${WHITE}in: ${RED}${UNTIL} ${WHITE}days
  ${RED}|${WHITE}|	  ${WHITE}Today ${RED}is: ${WHITE}${DATE}
  ${WHITE}|${RED}|	  ${RED}Gift ${WHITE}idea: ${RED}${GIFT}`)
	makeFile(filepath.Join(path, "themes", "candycane.theme"), candycane)

	json, err := json.MarshalIndent(defaultData, "", "")
	if err != nil {
		log.Fatal(err)
	}
	makeFile(filepath.Join(path, "config.json"), json)	
}


func makeFile(path string, data []byte) {
	if _, err := os.Stat(path); err != nil {
		_, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(path, data, 0644)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(colors.Green.Color + "[+] Created a file at " + colors.White.Color + path + colors.Reset.Color)
		}
	}
} 
