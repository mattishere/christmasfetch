package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)



type Config struct {
	Theme	string `json:"name"`
	Lights	string `json:"lights"`
}

func GetConfig() Config {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(configDir, "christmasfetch")
	
	if _, err := os.Stat(path); err != nil {
		err := os.Mkdir(path, 0700);
		if err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat(filepath.Join(path, "config.json")); err != nil {
		os.Create(filepath.Join(path, "config.json"))

		defaultData := Config{
			Theme: "gift",
			Lights: "random",
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

	return Config{}
}
