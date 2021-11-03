package logic

import "fmt"

const SIZEX = 9
const SIZEY = 15

func Hello() {
	fmt.Println("ggwp")
}

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

func (gol *GameOfLife) InitCells() {
	for index1, value1 := range gol.Cells {
		for index2, _ := range value1 {
			gol.Cells[index1][index2] = &Cell{
				PositionX: index1,
				PositionY: index2,
			}
		}
	}
}

func (gol *GameOfLife) markForChange(cell *Cell) {
	var x = cell.PositionX
	var y = cell.PositionY
	var sumOfNeighbours = 0

	var xlen = len(gol.Cells)
	var ylen = len(gol.Cells[0])
	// fmt.Println("x len", xlen)
	// fmt.Println("y len", ylen)
	if x > 0 && y > 0 {
		sumOfNeighbours += boolToInt(gol.Cells[x-1][y-1].IsAlive)
	}

	if x > 0 && y < ylen-1 {
		sumOfNeighbours += boolToInt(gol.Cells[x-1][y+1].IsAlive)
	}

	if x < xlen-1 && y > 0 {
		sumOfNeighbours += boolToInt(gol.Cells[x+1][y-1].IsAlive)
	}

	if x < xlen-1 && y < ylen-1 {
		sumOfNeighbours += boolToInt(gol.Cells[x+1][y+1].IsAlive)
	}

	if y < ylen-1 {
		sumOfNeighbours += boolToInt(gol.Cells[x][y+1].IsAlive)
	}

	if y > 0 {
		sumOfNeighbours += boolToInt(gol.Cells[x][y-1].IsAlive)
	}

	if x < xlen-1 {
		sumOfNeighbours += boolToInt(gol.Cells[x+1][y].IsAlive)
	}

	if x > 0 {
		sumOfNeighbours += boolToInt(gol.Cells[x-1][y].IsAlive)
	}

	if cell.IsAlive && (sumOfNeighbours < 2 || sumOfNeighbours > 3) {
		gol.Marked = append(gol.Marked, cell)
	} else if !cell.IsAlive && sumOfNeighbours == 3 {
		gol.Marked = append(gol.Marked, cell)
	}
}

func (gol *GameOfLife) SetCellState(x int, y int, IsAlive bool) {
	gol.Cells[x][y].IsAlive = IsAlive
}

func (gol *GameOfLife) Start(showCells func(*GameOfLife)) {
	for gol.Turn < 10 {
		gol.Turn++
		for _, value1 := range gol.Cells {
			for _, value2 := range value1 {
				gol.markForChange(value2)
			}
		}

		for _, val := range gol.Marked {
			val.IsAlive = !val.IsAlive
		}
		gol.Marked = []*Cell{}

		showCells(gol)
	}
}

func boolToInt(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
}
