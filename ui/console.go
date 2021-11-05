package ui

import (
	"fmt"
	"gameoflife/logic"
)

type ConsoleUI struct {
	Board *logic.GameOfLife
}

func (c *ConsoleUI) PrintBoard() {
	for _, value1 := range c.Board.Cells {
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

func (self *ConsoleUI) Play() {
	self.Board.SetCellState(4, 4, true)
	self.Board.SetCellState(4, 5, true)
	self.Board.SetCellState(4, 3, true)
	self.Board.SetCellState(5, 3, true)
	self.Board.SetCellState(5, 5, true)
	self.Board.SetCellState(6, 4, true)
	fmt.Println("START")
	fmt.Println("INITIAL FIELD:")
	self.PrintBoard()
	self.Board.Start(func(gol *logic.GameOfLife) {
		fmt.Println("TURN:", gol.Turn)
		self.PrintBoard()
	})

	fmt.Println("END")
}
