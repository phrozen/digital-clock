package tray

import (
	_ "embed"
	"runtime"

	"fyne.io/systray"
	"github.com/phrozen/digital-clock/clock/config"
)

//go:embed icon.ico
var iconICO []byte

//go:embed icon.png
var iconPNG []byte

type Tray struct {
	configFilePath string
}

func NewApp(configFilePath string) *Tray {
	cfg := config.NewConfigFromFile(configFilePath)
	cfg.WriteConfigToFile(configFilePath)
	return &Tray{
		configFilePath: configFilePath,
	}
}

func (t *Tray) OnExit() {

}

func (t *Tray) OnReady() {
	if runtime.GOOS == "windows" {
		systray.SetIcon(iconICO)
	} else {
		systray.SetIcon(iconPNG)
	}
	systray.SetTooltip("Digital Clock")
	mShow := systray.AddMenuItem("Hide", "Shows/Hides the digital clock")
	mReload := systray.AddMenuItem("Reload", "Reloads and updates clock from config file")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Exit", "Exit")

	// HANDLERS
	go func() {
		for {
			select {
			case <-mShow.ClickedCh:

			case <-mReload.ClickedCh:
				_ = config.NewConfigFromFile(t.configFilePath)
				// send new config to clock
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func (t *Tray) Run() error {
	// Fire up IPC

	systray.Run(t.OnReady, t.OnExit)
	return nil
}
