package config

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
		}
	}

	if _, err := os.Stat(filepath.Join(path, "config.json")); err != nil {
		_, err := os.Create(filepath.Join(path, "config.json"))
		if err != nil {
			log.Fatal(err)
		}

		jsonData, err := json.MarshalIndent(defaultData, "", "")
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(filepath.Join(path, "config.json"), jsonData, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(colors.Green.Color + "[+] New configuration file made at: " + colors.White.Color + filepath.Join(path, "config.json"))
}
