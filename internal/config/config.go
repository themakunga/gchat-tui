package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Theme struct {
	Colors struct {
		Backgroud  string `toml:"background"`
		Foreground string `toml:"foreground"`
		Primary    string `toml:"primary"`
		Secondary  string `toml:"secondary"`
		Success    string `toml:"success"`
		Warning    string `toml:"warning"`
		Error      string `toml:"error"`
	} `toml:"colors"`
}

type Config struct {
	Notifications struct {
		DesktopEnable bool `toml:"desktiop_enable"`
		SoundEnable   bool `toml:"sound_enable"`
		SoundFile     bool `toml:"sound_file"`
	} `toml:"notifications"`
}

type AppPaths struct {
	AuthDir   string
	ConfigDir string
	SoundDir  string
}

func SetupPaths() AppPaths {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not get home directory: %v", err)
	}

	configBase, err := os.UserConfigDir()
	if err != nil {
		configBase = filepath.Join(homeDir, ".config")
	}

	paths := AppPaths{
		AuthDir:   filepath.Join(homeDir, ".gchat-tui"),
		ConfigDir: filepath.Join(configBase, ".gchat-tui"),
	}
	paths.SoundDir = filepath.Join(paths.ConfigDir, "sounds")

	os.MkdirAll(paths.AuthDir, 0700)
	os.MkdirAll(paths.SoundDir, 0700)

	return paths
}

func LoadTheme(paths AppPaths) Theme {
	path := filepath.Join(paths.ConfigDir, "theme.toml")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		defaultTheme := `[colors]
background = "#1a1b26"
foreground = "#c0caf5"
primary    = "#7aa2f7"
secondary  = "#bb9af7"
success    = "#9ece6a"
warning    = "#e0af68"
error      = "#f7768e"
`
		os.WriteFile(path, []byte(defaultTheme), 0644)
	}

	var theme Theme
	b, _ := os.ReadFile(path)
	toml.Unmarshal(b, &theme)
	return theme
}

func LoadConfig(paths AppPaths) Config {
	path := filepath.Join(paths.ConfigDir, "config.toml")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		defaultConfig := `[notifications]
desktop_enabled = true
sound_enabled = true
sound_file = "sounds/gchat_alert.wav"
`
		os.WriteFile(path, []byte(defaultConfig), 0644)
	}

	var cfg Config
	b, _ := os.ReadFile(path)
	toml.Unmarshal(b, &cfg)
	return cfg
}
