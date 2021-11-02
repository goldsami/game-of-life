package main

import (
	"fmt"
	// "strconv"
	// "./gameoflife"
)

const SIZE = 9

type Cell struct {
	positionX int
	positionY int
	isAlive   bool
}

type GameOfLife struct {
	cells  [SIZE][SIZE]*Cell
	marked []*Cell
	turn   int
}

func (gol *GameOfLife) markForChange(cell *Cell) {
	var x = cell.positionX
	var y = cell.positionY
	var sumOfNeighbours = 0

	if x > 0 && y > 0 {
		sumOfNeighbours += boolToInt(gol.cells[x-1][y-1].isAlive)
	}

	if x > 0 && y < SIZE-1 {
		sumOfNeighbours += boolToInt(gol.cells[x-1][y+1].isAlive)
	}

	if x < SIZE-1 && y > 0 {
		sumOfNeighbours += boolToInt(gol.cells[x+1][y-1].isAlive)
	}

	if x < SIZE-1 && y < SIZE-1 {
		sumOfNeighbours += boolToInt(gol.cells[x+1][y+1].isAlive)
	}

	if y < SIZE-1 {
		sumOfNeighbours += boolToInt(gol.cells[x][y+1].isAlive)
	}

	if y > 0 {
		sumOfNeighbours += boolToInt(gol.cells[x][y-1].isAlive)
	}

	if x < SIZE-1 {
		sumOfNeighbours += boolToInt(gol.cells[x+1][y].isAlive)
	}

	if x > 0 {
		sumOfNeighbours += boolToInt(gol.cells[x-1][y].isAlive)
	}


	if cell.isAlive && (sumOfNeighbours < 2 || sumOfNeighbours > 3) {
		gol.marked = append(gol.marked, cell)
	} else if !cell.isAlive && sumOfNeighbours == 3 {
		gol.marked = append(gol.marked, cell)
	}
}

func (c *GameOfLife) setCellState(x int, y int, isAlive bool) {
	c.cells[x][y].isAlive = isAlive
}

func (gol *GameOfLife) start() {
	for gol.turn < 10 {
		gol.turn++
		fmt.Println("Turn:")
		fmt.Println(gol.turn)
		fmt.Println("========================")
		for _, value1 := range gol.cells {
			for _, value2 := range value1 {
				gol.markForChange(value2)
			}
		}

		for _, val := range gol.marked {
			val.isAlive = !val.isAlive
		}
		gol.marked = []*Cell{}
		printField(gol)
	}
}

func boolToInt(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
}

func printField(board *GameOfLife) {
	for _, value1 := range board.cells {
		var str string
		for _, value2 := range value1 {
			if value2.isAlive {
				str += "1"
			} else {
				str += "0"
			}
		}
		fmt.Println(str)
	}
}

func main() {
	var gol = GameOfLife{
		turn: 0,
	}

	for index1, value1 := range gol.cells {
		for index2, _ := range value1 {
			gol.cells[index1][index2] = &Cell{
				positionX: index1,
				positionY: index2,
			}
		}
	}

	gol.setCellState(4, 4, true)
	gol.setCellState(4, 5, true)
	gol.setCellState(4, 3, true)
	gol.setCellState(3, 4, true)
	gol.setCellState(5, 4, true)
	// gol.setCellState(4, 5, true)
	// gol.setCellState(4, 6, true)
	// gol.setCellState(4, 7, true)

	printField(&gol)
	// fmt.Println("111111111111111111111111111")
	gol.start()

}
