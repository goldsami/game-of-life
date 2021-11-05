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

func (gol *GameOfLife) InitEmptyCells() {
	for index1, value1 := range gol.Cells {
		for index2 := range value1 {
			gol.Cells[index1][index2] = &Cell{
				PositionX: index1,
				PositionY: index2,
			}
		}
	}
}

func (gol *GameOfLife) markCellForChange(cell *Cell) {
	var x = cell.PositionX
	var y = cell.PositionY
	var sumOfNeighbours = 0

	var xlen = len(gol.Cells)
	var ylen = len(gol.Cells[0])

	if x > 0 && y > 0 {
		sumOfNeighbours += utils.BoolToInt(gol.Cells[x-1][y-1].IsAlive)
	}

	if x > 0 && y < ylen-1 {
		sumOfNeighbours += utils.BoolToInt(gol.Cells[x-1][y+1].IsAlive)
	}

	if x < xlen-1 && y > 0 {
		sumOfNeighbours += utils.BoolToInt(gol.Cells[x+1][y-1].IsAlive)
	}

	if x < xlen-1 && y < ylen-1 {
		sumOfNeighbours += utils.BoolToInt(gol.Cells[x+1][y+1].IsAlive)
	}

	if y < ylen-1 {
		sumOfNeighbours += utils.BoolToInt(gol.Cells[x][y+1].IsAlive)
	}

	if y > 0 {
		sumOfNeighbours += utils.BoolToInt(gol.Cells[x][y-1].IsAlive)
	}

	if x < xlen-1 {
		sumOfNeighbours += utils.BoolToInt(gol.Cells[x+1][y].IsAlive)
	}

	if x > 0 {
		sumOfNeighbours += utils.BoolToInt(gol.Cells[x-1][y].IsAlive)
	}

	if cell.IsAlive && (sumOfNeighbours < 2 || sumOfNeighbours > 3) {
		gol.Marked = append(gol.Marked, cell)
	} else if !cell.IsAlive && sumOfNeighbours == 3 {
		gol.Marked = append(gol.Marked, cell)
	}
}

func (gol *GameOfLife) MarkForChange() {
	for _, value1 := range gol.Cells {
		for _, value2 := range value1 {
			gol.markCellForChange(value2)
		}
	}
	//fmt.Println("marked:", gol.Turn, ":", len(gol.Marked))
}

func (gol *GameOfLife) SetCellState(x int, y int, IsAlive bool) {
	//fmt.Println("set sate")
	gol.Cells[x][y].IsAlive = IsAlive
}

func (gol *GameOfLife) ResolveMarked() {
	for _, val := range gol.Marked {
		val.IsAlive = !val.IsAlive
	}
	gol.Marked = []*Cell{}
	//fmt.Println("Step")
}

func (gol *GameOfLife) Start(showCells func(*GameOfLife)) {
	for gol.Turn < 10 {
		gol.Turn++

		gol.MarkForChange()

		gol.ResolveMarked()

		//gol.Marked = []*Cell{}

		showCells(gol)

		//time.Sleep(2 * time.Second)
	}
}

func (gol *GameOfLife) MakeStep() {
	gol.MarkForChange()

	gol.ResolveMarked()

	//time.Sleep(2 * time.Second)
}
