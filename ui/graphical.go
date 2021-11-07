package ui

import (
	"fmt"
	"gameoflife/logic"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
)

const scale = 10

var black color.RGBA = color.RGBA{R: 75, G: 139, B: 190, A: 255}
var white color.RGBA = color.RGBA{R: 255, G: 232, B: 115, A: 255}
var counter int = 0

type GraphicalUI struct {
	Board *logic.GameOfLife
	Pause bool
}

func (self *GraphicalUI) Play() {
	self.Pause = true

	ebiten.Run(self.frame, len(self.Board.Cells)*scale, len(self.Board.Cells[0])*scale, 2, "Game of Life")
}

func (self *GraphicalUI) render(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()

	// Key pressed handlers
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		self.Board.Cells[mx/scale][my/scale].IsAlive = !self.Board.Cells[mx/scale][my/scale].IsAlive
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		fmt.Println("key pressed")
		self.Pause = !self.Pause
	}

	screen.Fill(white)

	// Fill screen with squares
	for x := 0; x < len(self.Board.Cells); x++ {
		for y := 0; y < len(self.Board.Cells[0]); y++ {
			if self.Board.Cells[x][y].IsAlive {
				for x1 := 0; x1 < scale; x1++ {
					for y1 := 0; y1 < scale; y1++ {
						screen.Set((x*scale)+x1, (y*scale)+y1, black)
					}
				}
			}
		}
	}
}

func (self *GraphicalUI) frame(screen *ebiten.Image) error {
	counter++
	var err error = nil
	if counter == 20 {
		if !self.Pause {
			self.Board.MakeStep()
		}
		counter = 0
	}

	if !ebiten.IsDrawingSkipped() {
		self.render(screen)
	}
	return err
}
