package clock

import (
	"image/color"
	"os"

	"github.com/BurntSushi/toml"
)

// Config is the configuration for Clock.
type Config struct {
	Seconds  bool
	Timezone bool
	UTC      bool
	Hours24  bool

	FontSize   float64
	FontColor  color.RGBA
	Background color.RGBA
}

// NewDefaultConfig returns a new Config with default values.
func NewDefaultConfig() *Config {
	config := &Config{
		FontSize:   64.0,
		FontColor:  color.RGBA{255, 255, 255, 255},
		Background: color.RGBA{0, 0, 0, 128},
		Seconds:    true,
		Timezone:   false,
		UTC:        false,
		Hours24:    false,
	}
	return config
}

func NewConfigFromFile(filename string) *Config {
	data, err := os.ReadFile(filename)
	if err != nil {
		return NewDefaultConfig()
	}
	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return NewDefaultConfig()
	}
	return &config
}

func (c *Config) WriteConfigToFile(filename string) error {
	data, err := toml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, os.ModePerm)
}
