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
		fmt.Println("Usage: game100 [X] [Y]")
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

	rand.NewSource(time.Now().UnixNano())

	var listOfMoves []string
	var gameTable = game.InitGameTable(10, "_")
	currValue := 1
	var currCellPtr = gameTable.GetCellAt(int8(startX), int8(startY))
	models.SetCellValue(currCellPtr, currValue) //starting point

	fmt.Println("Initial grid:")
	fmt.Println(gameTable)

	currCellPtr, movementDone, moveTries, err := game.DoNextMove(gameTable, currCellPtr, usedStrategy)
	totalTries := moveTries

	for {
		listOfMoves = append(listOfMoves, models.AllMovesNames[movementDone])
		currCellPtr, movementDone, moveTries, err = game.DoNextMove(gameTable, currCellPtr, usedStrategy)
		totalTries += moveTries
		if err.GameStatus == models.NoMoreMoves {
			break
		}

		//fmt.Println("Current Moves counter: ", len(listOfMoves))
	}

	fmt.Printf("Result: \n\tStarting point(x,y): [%d,%d]\n\tMoves: %d, \n\tTotal # of move tries: %d\n\tMax Value found: %s\n", startX, startY, len(listOfMoves), totalTries, currCellPtr.Val)
	fmt.Println(gameTable)
	fmt.Println(err)
	fmt.Println("Final list of Moves:", len(listOfMoves), ")\n", listOfMoves)
	if currCellPtr.Val == "100" {
		fmt.Println("Solution FOUND! I WON!!!!!!")
	}
	fmt.Println("END")

	//TODO: display list of rules
	//showRules()

}
