package clock

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Draggable is a draggable screen in Ebiten.
type Draggable struct {
	Count            int
	Dragging         bool
	DragStartCursorX int
	DragStartCursorY int
	CursorToWindowX  float64
	CursorToWindowY  float64
}

// Drag processess dragging logic to move the window poisition.
func (d *Draggable) Drag() {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		d.Dragging = false
	}
	if !d.Dragging && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		d.Dragging = true
		d.DragStartCursorX, d.DragStartCursorY = ebiten.CursorPosition()
	}
	if d.Dragging {
		cx, cy := ebiten.CursorPosition()
		dx := int(float64(cx-d.DragStartCursorX) * d.CursorToWindowX)
		dy := int(float64(cy-d.DragStartCursorY) * d.CursorToWindowY)
		wx, wy := ebiten.WindowPosition()
		ebiten.SetWindowPosition(wx+dx, wy+dy)
	}
}
