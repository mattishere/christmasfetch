package files

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Theme struct {
	Name string
	Art string
}

func GetThemes() []Theme {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(configDir, "christmasfetch", "themes")
	
	if _, err := os.Stat(path); err != nil {
		return []Theme{}
	}

	themes, err := os.ReadDir(path)
	if err != nil {
		return []Theme{}
	}

	var filtered []Theme
	for _, theme := range themes {
		if theme.IsDir() == false {
			name := theme.Name()
			extention := filepath.Ext(name)
			if extention == ".theme" {
				content, err := os.ReadFile(filepath.Join(path, name))
				if err != nil {
					log.Fatal(err)
				}
				final := Theme{
					Name: strings.ReplaceAll(name, extention, ""),
					Art: string(content),
				}
				filtered = append(filtered, final)
			}
		}
	}
	return filtered
}
