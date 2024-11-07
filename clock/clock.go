package clock

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed digital-7.mono.ttf
var font []byte
var fontFaceSource *text.GoTextFaceSource

func init() {
	// Create new font face source from TTF file.
	src, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		panic(err)
	}
	fontFaceSource = src
}

// Clock is a clock that draws the time in the given font.
type Clock struct {
	Draggable
	Config   *Config
	width    int
	height   int
	format   string
	font     *text.GoTextFace
	drawOps  *text.DrawOptions
	lastTime time.Time
	redraw   bool
}

// NewClock returns a new clock with the given configuration.
func NewClock(configuration *Config) *Clock {
	clk := &Clock{
		Config: configuration,
	}
	clk.LoadConfig()
	return clk
}

func (clk *Clock) LoadConfig() {
	clk.UpdateFontFace()
	clk.UpdateTimeFormat()
	clk.UpdateScreenSize()
	clk.redraw = true
}

func (clk *Clock) UpdateFontFace() {
	clk.font = &text.GoTextFace{
		Source: fontFaceSource,
		Size:   clk.Config.FontSize,
	}
	clk.drawOps = &text.DrawOptions{}
	// Add padding based on font size of the text.
	clk.drawOps.GeoM.Translate(clk.Config.FontSize/10, clk.Config.FontSize/20)
	clk.drawOps.ColorScale.ScaleWithColor(clk.Config.FontColor)
}

func (clk *Clock) UpdateTimeFormat() {
	clk.format = "03:04"
	if clk.Config.Hours24 {
		clk.format = "15:04"
	}
	if clk.Config.Seconds {
		clk.format += ":05"
	}
	if !clk.Config.Hours24 {
		clk.format += " PM"
	}
	if clk.Config.Timezone {
		clk.format += " MST"
	}
}

func (clk *Clock) UpdateScreenSize() {
	now := time.Now()
	if clk.Config.UTC {
		now = now.UTC()
	}
	current := now.Format(clk.format)
	width, height := text.Measure(current, clk.font, 0)
	clk.width = int(width+0.5) + int(clk.Config.FontSize/5.0)
	clk.height = int(height + 0.5)
	ebiten.SetWindowSize(clk.Width(), clk.Height())
}

// Widht returns the width of the clock.
func (clk *Clock) Width() int {
	return clk.width
}

// Height returns the height of the clock.
func (clk *Clock) Height() int {
	return clk.height
}

// ShouldRedraw returns true if the clock should be redrawn.
func (clk *Clock) ShouldRedraw(now time.Time) bool {
	if clk.redraw {
		return true
	}
	if clk.Config.Seconds && clk.lastTime.Second() != now.Second() {
		return true
	}
	if !clk.Config.Seconds && clk.lastTime.Minute() != now.Minute() {
		return true
	}
	return false
}

// Update checks for ebiten events and updates the clock.
func (clk *Clock) Update() error {
	// Ebiten automatically calls this function every frame
	clk.Drag()
	return nil
}

// Draw draws the clock on the screen.
func (clk *Clock) Draw(screen *ebiten.Image) {
	now := time.Now()
	// Only draw to the screen when neccessary to reduce CPU usage
	if !clk.ShouldRedraw(now) {
		return
	}
	clk.redraw = false
	// Clear the screen
	screen.Fill(clk.Config.Background)
	// IF UTC is true, convert to UTC and format accordingly
	if clk.Config.UTC {
		now = now.UTC()
	}
	// Draw the time on the screen
	text.Draw(screen, now.Format(clk.format), clk.font, clk.drawOps)
	clk.lastTime = time.Now()
}

// Layout returns the size of the clock.
func (clk *Clock) Layout(outsideWidth, outsideHeight int) (int, int) {
	clk.Draggable.CursorToWindowX = float64(outsideWidth) / float64(clk.width)
	clk.Draggable.CursorToWindowY = float64(outsideHeight) / float64(clk.height)
	return clk.width, clk.height
}
