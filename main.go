package main

import (
	"fmt"
	"game-of-100/game"
	"game-of-100/models"
	"game-of-100/strategies"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if !(len(os.Args) == 5 || len(os.Args) == 3) {
		fmt.Println("Usage: game100 STRATEGY MAX_TIMEOUT_SEC [X] [Y]")
		os.Exit(1)
	}

	startX := 0
	startY := 0

	usedStrategy := readStrategy()
	timeout := readTimeout()

	if len(os.Args) == 5 {
		startX, startY = readStartPosition()
	}

	fmt.Println("Ready to find a solution to the Game Of 100 puzzle!")
	fmt.Println("Max Timeout: ", timeout.String())
	fmt.Println("Chosen strategy: ", strategies.AllNames[usedStrategy])

	rand.NewSource(time.Now().UnixNano())
	maxValueFound := 1
	startTime := time.Now()

	var listOfMoves []string
	var err models.StatusErr
	var totalTries int
	var gameTable models.Table
	var currCellPtr *models.Cell

	fmt.Println("Start!")
	var executionCounter = 0
	var maxValueTable models.Table
	var maxValueMoves []string
	const gridSize = 10
	const fullGridValue = gridSize * gridSize

	for {
		gameTable = game.InitGameTable(gridSize, "_")
		currValue := 1
		currCellPtr = gameTable.GetCellAt(int8(startX), int8(startY))
		models.SetCellValue(currCellPtr, currValue)
		executionTries := 0

		if maxValueFound == 1 {
			fmt.Println("Initial grid:")
			fmt.Println(gameTable)
			fmt.Println("Value to reach: ", fullGridValue)
		}

		listOfMoves, gameTable, currCellPtr, executionTries, err = GameExecution(gameTable, currCellPtr, usedStrategy)

		valFound, _ := strconv.Atoi(currCellPtr.Val)
		executionCounter += 1
		totalTries += executionTries

		if valFound > maxValueFound {
			maxValueFound = valFound
			fmt.Printf("New MaxValue %d found after %d executions and %d tries\n", maxValueFound, executionCounter, executionTries)
			maxValueTable = gameTable
			maxValueMoves = listOfMoves
		}

		if maxValueFound == fullGridValue {
			break
		}

		elapsed := time.Since(startTime)
		if elapsed > timeout {
			fmt.Println("Timeout reached")
			break
		}
	}
	fmt.Println(maxValueTable)
	fmt.Println(err)
	fmt.Println("Moves:", maxValueMoves)
	if currCellPtr.Val == strconv.Itoa(fullGridValue) {
		fmt.Println("Solution FOUND! I WON!!!!!!")
		fmt.Println("Found in: ", time.Since(startTime).Milliseconds(), " msec")
	}
	fmt.Println("Executions number and total number of move tries: ", executionCounter, totalTries)
	fmt.Println("END")

}

func GameExecution(gameTable models.Table, currCellPtr *models.Cell, usedStrategy strategies.Strategy) ([]string, models.Table, *models.Cell, int, models.StatusErr) {
	var listOfMoves []string

	currCellPtr, movementDone, moveTries, err := game.DoNextMove(gameTable, currCellPtr, usedStrategy)
	totalTries := moveTries

	for {
		listOfMoves = append(listOfMoves, models.AllMovesNames[movementDone])
		currCellPtr, movementDone, moveTries, err = game.DoNextMove(gameTable, currCellPtr, usedStrategy)
		totalTries += moveTries
		if err.GameStatus == models.NoMoreMoves {
			break
		}

	}
	return listOfMoves, gameTable, currCellPtr, totalTries, err
}
