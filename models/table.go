package models

import (
	"fmt"
	"strconv"
)

const emptyCellValue string = "_"

type TableBehaviour interface {
	Move(*Cell, Movement) (*Cell, error)
}

type Table struct {
	lines [][]*Cell
}

func MakeTable(lines [][]*Cell) Table {
	return Table{lines: lines}
}

func (g Table) String() string {
	result := ""
	for _, row := range g.lines {
		for _, cell := range row {
			result += fmt.Sprintf("|%3s| ", cell.Val)
		}
		result += "\n"
	}
	return result
}

func (t Table) GetCellAt(x, y int8) *Cell {
	return t.lines[y][x]
}

func (t Table) Move(currCellPtr *Cell, movement Movement) (*Cell, error) {
	gridSize := t.GridSize()

	nextX, nextY := GetNextPositionByMove(movement, currCellPtr.xPos, currCellPtr.yPos)

	if IsOutOfBound(nextX, nextY, gridSize) {
		return nil, StatusErr{
			GameStatus: InvalidPosition,
			Message:    fmt.Sprintf("[%d, %d] is out of grid bounds", nextX, nextY),
		}
	}

	var currValue, err = strconv.Atoi(currCellPtr.Val)

	if err != nil {
		return nil, StatusErr{
			GameStatus: InvalidCellValue,
			Message:    fmt.Sprintf("[%s] invalid current cell value", currCellPtr.Val),
		}
	}

	nextValue := currValue + 1

	nextCellPtr := t.lines[nextY][nextX]

	if nextCellPtr.Val == emptyCellValue {
		// empty cell, we can return it
		SetCellValue(nextCellPtr, nextValue)
		return nextCellPtr, nil
	}

	return nil, StatusErr{
		GameStatus: CellAlreadyUsed,
		Message:    fmt.Sprintf("[%d, %d] is not empty. Can't proceed", nextX, nextY),
	}
}

func (t Table) GridSize() int8 {
	return int8(len(t.lines[0]))
}

func IsOutOfBound(x int8, y int8, gridSize int8) bool {
	isLeftRightOut := y >= gridSize || y < 0
	isTopBottomOut := x >= gridSize || x < 0
	return isLeftRightOut || isTopBottomOut
}

func SetCellValue(gameCell *Cell, initValue int) {
	gameCell.Val = strconv.Itoa(initValue)
}

func GetNextPositionByMove(dir Movement, x, y int8) (int8, int8) {
	var nextX, nextY int8

	//fmt.Printf("From position: [%d,%d]\n", x, y)
	switch dir {
	case Up:
		//	fmt.Println("Move up")
		nextX = x
		nextY = y - 3
	case Down:
		//	fmt.Println("Move down")
		nextX = x
		nextY = y + 3
	case Left:
		//	fmt.Println("Move left")
		nextX = x - 3
		nextY = y
	case Right:
		//	fmt.Println("Move right")
		nextX = x + 3
		nextY = y
	case UpLeft:
		//	fmt.Println("Move diagonally up-left")
		nextX = x - 2
		nextY = y - 2
	case UpRight:
		//	fmt.Println("Move diagonally up-right")
		nextX = x + 2
		nextY = y - 2
	case DownLeft:
		//	fmt.Println("Move diagonally down-left")
		nextX = x - 2
		nextY = y + 2
	case DownRight:
		//	fmt.Println("Move diagonally down-right")
		nextX = x + 2
		nextY = y + 2
	default:
		panic("Unknown direction")
	}

	//fmt.Printf("New position: [%d,%d]\n\n", nextX, nextY)

	return nextX, nextY
}
