package tray

import (
	_ "embed"
	"log"

	"fyne.io/systray"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/phrozen/digital-clock/clock"
)

//go:embed icon.ico
var icon []byte

type Tray struct {
	clk            *clock.Clock
	configFilePath string
}

func NewApp(configFilePath string) *Tray {
	cfg := clock.NewConfigFromFile(configFilePath)
	cfg.WriteConfigToFile(configFilePath)
	return &Tray{
		clk:            clock.NewClock(cfg),
		configFilePath: configFilePath,
	}
}

func (t *Tray) OnExit() {
	log.Println("tray exit")
}

func (t *Tray) OnReady() {
	systray.SetIcon(icon)
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
				if ebiten.IsWindowMinimized() {
					ebiten.RestoreWindow()
					mShow.SetTitle("Hide")
				} else {
					ebiten.MinimizeWindow()
					mShow.SetTitle("Show")
				}
			case <-mReload.ClickedCh:
				cfg := clock.NewConfigFromFile(t.configFilePath)
				t.clk.Config = cfg
				t.clk.LoadConfig()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()

	// Run ebiten app with configured clock
	ebiten.SetWindowSize(t.clk.Width(), t.clk.Height())
	ebiten.SetWindowTitle("Digital Clock")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowDecorated(false)
	ebiten.SetScreenClearedEveryFrame(false)
	op := &ebiten.RunGameOptions{}
	op.ScreenTransparent = true
	op.SkipTaskbar = true
	if err := ebiten.RunGameWithOptions(t.clk, op); err != nil {
		log.Println("error:", err)
	}
}
