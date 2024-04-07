package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const emptyCellValue string = "_"

type GameStatus int

const (
	InvalidPosition GameStatus = iota + 1
	InvalidCellValue
	CellAlreadyUsed
	NoMoreMoves
)

type GameStatusErr struct {
	GameStatus GameStatus
	Message    string
}

func (gse GameStatusErr) Error() string {
	return gse.Message
}

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

var allMoves = []Movement{Up, Down, Left, Right, DownRight, DownLeft, UpLeft, UpRight}

var moveNames = [...]string{
	"Up",
	"Down",
	"Left",
	"Right",
	"UpLeft",
	"UpRight",
	"DownLeft",
	"DownRight",
}

type GameCell struct {
	val  string
	xPos int8
	yPos int8
}

type GameTable struct {
	lines [][]*GameCell
}

func (g GameTable) move(currCellPtr *GameCell, movement Movement) (*GameCell, error) {
	gridSize := int8(len(g.lines[0]))

	nextX, nextY := getNextPositionByMove(movement, currCellPtr.xPos, currCellPtr.yPos)

	if isOutOfGridBound(nextX, nextY, gridSize) {
		return nil, GameStatusErr{
			GameStatus: InvalidPosition,
			Message:    fmt.Sprintf("[%d, %d] is out of grid bounds", nextX, nextY),
		}
	}

	var currValue, err = strconv.Atoi(currCellPtr.val)

	if err != nil {
		return nil, GameStatusErr{
			GameStatus: InvalidCellValue,
			Message:    fmt.Sprintf("[%s] invalid current cell value", currCellPtr.val),
		}
	}

	nextValue := currValue + 1

	nextCellPtr := g.lines[nextY][nextX]

	if nextCellPtr.val == emptyCellValue {
		// empty cell, we can return it
		setCellValue(nextCellPtr, nextValue)
		return nextCellPtr, nil
	}

	return nil, GameStatusErr{
		GameStatus: CellAlreadyUsed,
		Message:    fmt.Sprintf("[%d, %d] is not empty. Can't proceed", nextX, nextY),
	}
}

func getNextPositionByMove(dir Movement, x, y int8) (int8, int8) {
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

func isOutOfGridBound(x int8, y int8, gridSize int8) bool {
	isLeftRightOut := y >= gridSize || y < 0
	isTopBottomOut := x >= gridSize || x < 0
	return isLeftRightOut || isTopBottomOut
}

func (g GameTable) String() string {
	result := ""
	for _, row := range g.lines {
		for _, cell := range row {
			result += fmt.Sprintf("|%2s| ", cell.val)
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
			line = append(line, makeCellPtr(defaultVal, int8(i), int8(r)))
		}
		lines = append(lines, line)
	}

	return GameTable{lines: lines}
}

func setCellValue(gameCell *GameCell, initValue int) {
	gameCell.val = strconv.Itoa(initValue)
}

func main() {
	if !(len(os.Args) == 3 || len(os.Args) == 1) {
		fmt.Println("Usage: game100 [X] [Y]")
		os.Exit(1)
	}

	startX := 0
	startY := 0
	var paramsErr []error = make([]error, 2)

	if len(os.Args) == 3 {
		startX, paramsErr[0] = strconv.Atoi(os.Args[1])
		startY, paramsErr[1] = strconv.Atoi(os.Args[2])
		if paramsErr[0] != nil || paramsErr[1] != nil || isOutOfGridBound(int8(startX), int8(startY), 10) {
			fmt.Println("invalid starting point: [", os.Args[1], ",", os.Args[2], "]")
			os.Exit(1)
		}
	}

	fmt.Println("Ready to find a solution to the Game Of 100 solitaire!")

	rand.NewSource(time.Now().UnixNano())

	var listOfMoves []string
	var gameTable = initGameTable(10, "_")
	currValue := 1
	var currCellPtr = gameTable.lines[startY][startX]
	setCellValue(currCellPtr, currValue) //starting point

	fmt.Println("Initial grid:")
	fmt.Println(gameTable)

	currCellPtr, movementDone, moveTries, err := takeNextMove(gameTable, currCellPtr)
	totalTries := moveTries

	for {
		listOfMoves = append(listOfMoves, moveNames[movementDone])
		currCellPtr, movementDone, moveTries, err = takeNextMove(gameTable, currCellPtr)
		totalTries += moveTries
		if err.GameStatus == NoMoreMoves {
			break
		}

		//fmt.Println("Current Moves counter: ", len(listOfMoves))
	}

	fmt.Printf("Result: \n\tStarting point(x,y): [%d,%d]\n\tMoves: %d, \n\tTotal # of move tries: %d\n\tMax Value found: %s\n", startX, startY, len(listOfMoves), totalTries, currCellPtr.val)
	fmt.Println(gameTable)
	fmt.Println(err)
	fmt.Println("Final list of Moves:", len(listOfMoves), ")\n", listOfMoves)
	fmt.Println("END")

	//TODO: display list of rules
	//showRules()

}

func takeNextMove(gameTable GameTable, fromCellPtr *GameCell) (*GameCell, Movement, int, GameStatusErr) {

	allMovesMap := make(map[Movement]bool)
	for _, dir := range allMoves {
		allMovesMap[dir] = false
	}

	var currCellPtr = fromCellPtr
	var nextMovePtr *Movement
	var err GameStatusErr

	nextMovePtr, err = getNextFreeMove(allMovesMap)
	tries := 1

	for err.GameStatus != NoMoreMoves {
		//fmt.Println("Tentative Move", tries)
		currCellPtr, moveErr := gameTable.move(currCellPtr, *nextMovePtr)
		if moveErr == nil {
			//found a valid move, return it
			return currCellPtr, *nextMovePtr, tries, GameStatusErr{}
		}
		allMovesMap[*nextMovePtr] = true
		tries += 1
		nextMovePtr, err = getNextFreeMove(allMovesMap)
	}

	return fromCellPtr, -1, tries, GameStatusErr{GameStatus: NoMoreMoves, Message: "No more free moves"}
}

func getNextFreeMove(moves map[Movement]bool) (*Movement, GameStatusErr) {
	var validMoves []Movement
	for _, m := range allMoves {
		if !moves[m] {
			validMoves = append(validMoves, m)
		}
	}

	sizeValidMoves := len(validMoves)
	if sizeValidMoves != 0 {
		validMove := validMoves[rand.Intn(sizeValidMoves)]
		return &validMove, GameStatusErr{}
	}

	return nil, GameStatusErr{GameStatus: NoMoreMoves, Message: "No more free moves"}

}
