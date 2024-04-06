package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const emptyCellValue string = "_"

type Movement int

const (
	Up Movement = iota
	Down
	Left
	Right
	UpLeft
	UpRight
	DownLeft
	DownRight
)

type GameCell struct {
	val  string
	xPos int8
	yPos int8
	// HorizontalMovement
}

type GameTable struct {
	lines [][]*GameCell
}

func (g GameTable) move(currCellPtr *GameCell, movement Movement) (*GameCell, error) {
	gridSize := int8(len(g.lines[0]))

	nextX, nextY := getNextPositionByMove(movement, currCellPtr.xPos, currCellPtr.yPos)

	if isOutOfGridBound(nextX, nextY, gridSize) {
		return &GameCell{}, fmt.Errorf("[%d, %d] is out of grid bounds", nextX, nextY)
	}

	var currValue, err = strconv.Atoi(currCellPtr.val)

	if err != nil {
		return &GameCell{}, fmt.Errorf("[%s] invalid current cell value", currCellPtr.val)
	}

	nextValue := currValue + 1

	nextCellPtr := g.lines[nextX][nextY]

	if nextCellPtr.val == emptyCellValue {
		// empty cell, we can return it
		setCellValue(nextCellPtr, nextValue)
		return nextCellPtr, nil
	}

	return &GameCell{}, fmt.Errorf("[%d, %d] is not empty. Can't proceed", nextX, nextY)
}

func getNextPositionByMove(dir Movement, x, y int8) (int8, int8) {
	var nextX, nextY int8

	switch dir {
	case Up:
		fmt.Println("Move up")
		nextX = x
		nextY = y - 3
	case Down:
		fmt.Println("Move down")
		nextX = x
		nextY = y + 3
	case Left:
		fmt.Println("Move left")
		nextX = x - 3
		nextY = y
	case Right:
		fmt.Println("Move right")
		nextX = x + 3
		nextY = y
	case UpLeft:
		fmt.Println("Move diagonally up-left")
		nextX = x - 2
		nextY = y - 2
	case UpRight:
		fmt.Println("Move diagonally up-right")
		nextX = x + 2
		nextY = y - 2
	case DownLeft:
		fmt.Println("Move diagonally down-left")
		nextX = x - 2
		nextY = y + 2
	case DownRight:
		fmt.Println("Move diagonally down-right")
		nextX = x + 2
		nextY = y + 2
	default:
		panic("Unknown direction")
	}

	return nextX, nextY

}

func isOutOfGridBound(x int8, y int8, gridSize int8) bool {
	isLeftRightOut := y > gridSize || y < 0
	isTopBottomOut := x > gridSize || x < 0
	//TODO: verify check 
	return isLeftRightOut || isTopBottomOut
}

func (g GameTable) String() string {
	result := ""
	for _, row := range g.lines {
		for _, cell := range row {
			result += fmt.Sprintf("|%s| ", cell.val)
		}
		result += "\n"
	}
	return result
}

func makeCellPtr(val string, xPos, yPos int8) *GameCell {
	return &GameCell{val: val, xPos: xPos, yPos: yPos}
}

func initGameTable(n int, defaultVal string) GameTable {
	var lines [][]*GameCell

	for r := 0; r < n; r++ {
		var line []*GameCell
		for i := 0; i < n; i++ {
			line = append(line, makeCellPtr(defaultVal, int8(r), int8(i)))
		}
		lines = append(lines, line)
	}

	return GameTable{lines: lines}
}

func setCellValue(gameCell *GameCell, initValue int) {
	gameCell.val = strconv.Itoa(initValue)
}

func main() {
	fmt.Println("Ready to find a solution to the Game Of 100 solitaire!")

	rand.NewSource(time.Now().UnixNano())

	// Generate a random number between 0 and 7 (inclusive) to represent a direction
	randomMovement := Movement(rand.Intn(8))

	var gameTable = initGameTable(10, "_")

	fmt.Println("Initial grid:")
	//fmt.Println(grid)

	currValue := 1
	var currCellPtr = gameTable.lines[5][5]

	setCellValue(currCellPtr, currValue) //starting point

	fmt.Println(gameTable)
	currValue += 1

	currCell, err := gameTable.move(currCellPtr, randomMovement)

	for err == nil {
		//setCellValue(currCell, currValue)
		fmt.Println(gameTable)
		randomMovement = Movement(rand.Intn(8))
		currCell, err = gameTable.move(currCell, randomMovement)
		//currValue += 1
	}

	fmt.Println(err)
	fmt.Println("END")

	//TODO: display list of rules
	//showRules()

}
