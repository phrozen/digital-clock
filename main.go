package main

import (
	_ "embed"
	"flag"

	"github.com/getlantern/systray"
	"github.com/phrozen/digital-clock/tray"
)

func main() {
	configFilePath := flag.String("config", "config.toml", "Path to the configuration file to be used")
	flag.Parse()

	app := tray.NewApp(*configFilePath)
	systray.Run(app.OnReady, app.OnExit)
}
