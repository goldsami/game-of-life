package main

import (
	"bufio"
	"fmt"
	"gameoflife/logic"
	"gameoflife/ui"
	"os"
	"strings"
)

func main() {
	var gol = logic.GameOfLife{
		Turn: 0,
	}
	gol.InitEmptyCells()

	var userinterface ui.UI

	fmt.Println("Game of Life")
	for true {
		fmt.Println("---------------------")
		fmt.Println("Press '1' to play in console.\n Press '2' to play in GUI. \n Press '0' to exit")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		result := strings.TrimRight(text, "\n\r")

		if result == "0" {
			break
		} else if result == "1" {
			userinterface = &ui.ConsoleUI{
				Board: &gol,
			}
			break
		} else if result == "2" {
			userinterface = &ui.GraphicalUI{
				Board: &gol,
			}
			break
		}
	}

	userinterface.Play()
}
