package ui

import (
	"fmt"
	"gameoflife/logic"
)

func PrintBoard(board *logic.GameOfLife) {
	for _, value1 := range board.Cells {
		var str string
		for _, value2 := range value1 {
			if value2.IsAlive {
				str += "1"
			} else {
				str += "0"
			}
		}
		fmt.Println(str)
	}
}

func Play(gol *logic.GameOfLife) {
	gol.SetCellState(4, 4, true)
	gol.SetCellState(4, 5, true)
	gol.SetCellState(4, 3, true)
	gol.SetCellState(3, 4, true)
	gol.SetCellState(5, 4, true)
	fmt.Println("START")
	fmt.Println("INITIAL FIELD:")
	PrintBoard(gol)
	gol.Start(func(gol *logic.GameOfLife) {
		fmt.Println("TURN:", gol.Turn)
		PrintBoard(gol)
	})
	fmt.Println("END")
}
