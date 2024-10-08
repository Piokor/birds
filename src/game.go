package birds

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	flock Flock
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	target := Vertex{float64(x), float64(y)}
	g.flock.Update(target)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.flock.Draw(screen)
}

func NewGame(flock Flock) *Game {
	return &Game{flock}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
