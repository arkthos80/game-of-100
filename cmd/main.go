package main

import "fmt"
import "github.com/arkthos80/game-of-100/gameof100"

func main() {
	fmt.Println("Ready to find a solution to the Game Of 100 solitaire!")

	grid := Grid{}

	for _, l := range grid.lines {
		for r := range l {
			fmt.Print(r, ", ")
		}
		fmt.Println()
	}
}
