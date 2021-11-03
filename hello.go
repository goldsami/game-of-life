package main

import (
	// "fmt"
	"gameoflife/ui"
	"gameoflife/logic"
)

func main() {
	var gol = logic.GameOfLife{
		Turn: 0,
	}

	gol.InitCells()

	ui.Play(&gol)
}
