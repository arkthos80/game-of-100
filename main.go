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
	if !(len(os.Args) == 3 || len(os.Args) == 1) {
		fmt.Println("Usage: game100 [X] [Y] ")
		os.Exit(1)
	}

	startX := 0
	startY := 0
	usedStrategy := strategies.RandomHVFirstDiagRandom
	var paramsErr []error = make([]error, 2)

	if len(os.Args) == 3 {
		startX, paramsErr[0] = strconv.Atoi(os.Args[1])
		startY, paramsErr[1] = strconv.Atoi(os.Args[2])
		if paramsErr[0] != nil || paramsErr[1] != nil || models.IsOutOfBound(int8(startX), int8(startY), 10) {
			fmt.Println("invalid starting point: [", os.Args[1], ",", os.Args[2], "]")
			os.Exit(1)
		}
	}

	fmt.Println("Ready to find a solution to the Game Of 100 solitaire!")
	timeout := 10 * time.Second
	fmt.Println("Max Timeout: " + timeout.String() + " seconds")

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

	for {
		gameTable = game.InitGameTable(10, "_")
		currValue := 1
		currCellPtr = gameTable.GetCellAt(int8(startX), int8(startY))
		models.SetCellValue(currCellPtr, currValue)
		executionTries := 0

		if maxValueFound == 1 {
			fmt.Println("Initial grid:")
			fmt.Println(gameTable)
		}

		listOfMoves, gameTable, currCellPtr, executionTries, err = GameExecution(gameTable, currCellPtr, usedStrategy)

		valFound, _ := strconv.Atoi(currCellPtr.Val)
		executionCounter += 1
		totalTries += executionTries

		if valFound > maxValueFound {
			maxValueFound = valFound
			fmt.Println("New MaxValue found: ", maxValueFound)
			maxValueTable = gameTable
			maxValueMoves = listOfMoves
		}

		if maxValueFound == 100 {
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
	if currCellPtr.Val == "100" {
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
