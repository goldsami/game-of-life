package main

import (
	"gameoflife/logic"
	"gameoflife/ui"
)

func main() {
	var gol = logic.GameOfLife{
		Turn: 0,
	}
	gol.InitEmptyCells()

	ui.Play(&gol)
}
