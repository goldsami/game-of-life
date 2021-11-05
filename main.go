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

	//var userinterface ui.UI = &ui.ConsoleUI{
	//	Board: &gol,
	//}
	var userinterface ui.UI = &ui.GraphicalUI{
		Board: &gol,
	}

	userinterface.Play()
}
