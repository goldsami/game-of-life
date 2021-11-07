package main

import (
	"gameoflife/logic"
	"testing"
)

func TestMakeOneStep(t *testing.T) {
	var gol = logic.GameOfLife{
		Turn: 0,
	}
	gol.InitEmptyCells()

	gol.SetCellState(1, 1, true)
	gol.SetCellState(1, 2, true)
	gol.SetCellState(1, 3, true)

	gol.MakeStep()

	if gol.Cells[1][1].IsAlive || !gol.Cells[1][2].IsAlive || gol.Cells[1][3].IsAlive || !gol.Cells[0][2].IsAlive || !gol.Cells[2][2].IsAlive {
		t.Fatalf("Wrong field state")
	}
}
