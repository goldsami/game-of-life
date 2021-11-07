package logic

import (
	"gameoflife/utils"
)

const SIZEX = 25
const SIZEY = 25

type Cell struct {
	PositionX int
	PositionY int
	IsAlive   bool
}

type GameOfLife struct {
	Cells  [SIZEX][SIZEY]*Cell
	Marked []*Cell
	Turn   int
}

func (self *GameOfLife) InitEmptyCells() {
	for index1, value1 := range self.Cells {
		for index2 := range value1 {
			self.Cells[index1][index2] = &Cell{
				PositionX: index1,
				PositionY: index2,
			}
		}
	}
}

func (self *GameOfLife) markCellForChange(cell *Cell) {
	var x = cell.PositionX
	var y = cell.PositionY
	var sumOfNeighbours = 0

	var xlen = len(self.Cells)
	var ylen = len(self.Cells[0])

	// Calculate sum of neighbours
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if !(i == x && j == y) && i >= 0 && j >= 0 && i <= xlen-1 && j <= ylen-1 {
				sumOfNeighbours += utils.BoolToInt(self.Cells[i][j].IsAlive)
			}
		}
	}

	// Mark cell to be changed next turn
	if cell.IsAlive && (sumOfNeighbours < 2 || sumOfNeighbours > 3) {
		self.Marked = append(self.Marked, cell)
	} else if !cell.IsAlive && sumOfNeighbours == 3 {
		self.Marked = append(self.Marked, cell)
	}
}

func (self *GameOfLife) MarkForChange() {
	for _, value1 := range self.Cells {
		for _, value2 := range value1 {
			self.markCellForChange(value2)
		}
	}
}

func (self *GameOfLife) SetCellState(x int, y int, IsAlive bool) {
	self.Cells[x][y].IsAlive = IsAlive
}

func (self *GameOfLife) ResolveMarked() {
	for _, val := range self.Marked {
		val.IsAlive = !val.IsAlive
	}
	self.Marked = []*Cell{}
}

func (self *GameOfLife) Start(showCells func(*GameOfLife)) {
	for self.Turn < 10 {
		self.Turn++

		self.MakeStep()

		showCells(self)
	}
}

func (self *GameOfLife) MakeStep() {
	self.MarkForChange()
	self.ResolveMarked()
}
