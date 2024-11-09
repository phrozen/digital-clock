package main

import (
	_ "embed"
	"flag"
	"log"

	"github.com/phrozen/digital-clock/clock"
	"github.com/phrozen/digital-clock/tray"
)

type App interface {
	Run() error
}

func main() {
	configFilePath := flag.String("config", "config.toml", "Path to the configuration file to be used")
	mode := flag.String("mode", "tray", "Mode to run the app in (tray, clock)")
	flag.Parse()

	var app App
	switch *mode {
	case "tray":
		app = tray.NewApp(*configFilePath)
	case "clock":
		app = clock.NewClock(*configFilePath)
	default:
		panic("Invalid mode")
	}

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
