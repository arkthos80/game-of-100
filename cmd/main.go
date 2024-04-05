package main

import (
	"fmt"
	"strconv"
)

//TODO: use pointers to GameCell

type HorizontalMovement interface {
	MoveLeft() GameCell
	MoveRight() GameCell
}

type VerticalMovement interface {
	MoveTop() GameCell
	MoveBottom() GameCell
}

type DiagonalMovement interface {
	MoveTopLeft() GameCell
	MoveTopRight() GameCell
	MoveBottomLeft() GameCell
	MoveBottomRight() GameCell
}

type GameCell struct {
	val  string
	xPos int8
	yPos int8
	HorizontalMovement
}

type Grid struct {
	lines [][]*GameCell
}

func (g Grid) moveHorizontal(currCell GameCell, left bool) (GameCell, error) {
	gridSize := int8(len(g.lines))
	var nextYpos int8

	if left {
		nextYpos = currCell.yPos - 2
	} else {
		nextYpos = currCell.yPos + 2
	}

	var out, bound = isOutOfGridBound(nextYpos, gridSize, left)

	if out {
		return GameCell{}, fmt.Errorf("%d is out of %s grid bound", nextYpos, bound)
	}

	nextCell := g.lines[currCell.xPos][nextYpos]

	if nextCell.val == "_" { //call a method like isEmptyCell
		// empty cell, we can return it
		return nextCell, nil
	}

	return GameCell{}, fmt.Errorf("%d is not empty. Can't proceed", nextYpos)
}

func isOutOfGridBound(nextYPos int8, gridSize int8, leftSide bool) (bool, string) {
	var bound = "right"
	if leftSide {
		bound = "left"
	}

	return nextYPos > gridSize || nextYPos < 0, bound
}

func (g Grid) String() string {
	result := ""
	for _, row := range g.lines {
		for _, cell := range row {
			result += fmt.Sprintf("|%s| ", cell.val)
		}
		result += "\n"
	}
	return result
}

func makeCell(val string, xPos, yPos int8) GameCell {
	return GameCell{val: val, xPos: xPos, yPos: yPos}
}

func makeGrid(n int, defaultVal string) Grid {
	var lines [][]GameCell

	for r := 0; r < n; r++ {
		var line []GameCell
		for i := 0; i < n; i++ {
			line = append(line, makeCell(defaultVal, int8(r), int8(i)))
		}
		lines = append(lines, line)
	}

	return Grid{lines: lines}
}

func setCellValue(gameCell *GameCell, initValue int) {
	gameCell.val = strconv.Itoa(initValue)
}

func main() {
	fmt.Println("Ready to find a solution to the Game Of 100 solitaire!")

	var grid = makeGrid(10, "_")

	fmt.Println("Initial grid:")
	//fmt.Println(grid)

	currValue := 1
	setCellValue(&grid.lines[0][0], currValue) //starting point

	fmt.Println(grid)
	currValue += 1

	currCell, err := grid.moveHorizontal(grid.lines[0][0], false)

	for err == nil {
		setCellValue(&currCell, currValue)
		fmt.Println(grid)
		currCell, err = grid.moveHorizontal(currCell, false)
		currValue += 1
	}

	fmt.Println(err)
	fmt.Println("END")

	//TODO: display list of rules
	//showRules()

}
