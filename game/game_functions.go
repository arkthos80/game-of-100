package game

import (
	"game-of-100/models"
	"game-of-100/strategies"
	"math/rand"
)

func DoNextMove(gameTable models.Table, fromCellPtr *models.Cell, strategy strategies.Strategy) (*models.Cell, models.Movement, int, models.StatusErr) {

	allMovesMap := make(map[models.Movement]bool)
	for _, dir := range models.AllMoves {
		allMovesMap[dir] = false
	}

	var currCellPtr = fromCellPtr
	var nextMovePtr *models.Movement
	var err models.StatusErr

	nextMovePtr, err = getNextMoveByStrategy(allMovesMap, strategy)
	tries := 1

	for err.GameStatus != models.NoMoreMoves {
		//fmt.Println("Tentative Move", tries)
		currCellPtr, moveErr := gameTable.Move(currCellPtr, *nextMovePtr)
		if moveErr == nil {
			//found a valid move, return it
			return currCellPtr, *nextMovePtr, tries, models.StatusErr{}
		}
		allMovesMap[*nextMovePtr] = true
		tries += 1
		nextMovePtr, err = getNextMoveByStrategy(allMovesMap, strategy)
	}

	return fromCellPtr, -1, tries, models.StatusErr{GameStatus: models.NoMoreMoves, Message: "No more free moves"}
}

func getNextMoveByStrategy(moves map[models.Movement]bool, strategy strategies.Strategy) (*models.Movement, models.StatusErr) {
	validMoves, sizeValidMoves := getRemainingMoves(moves, models.AllMoves)

	if sizeValidMoves != 0 {
		var nextMove models.Movement
		switch strategy {
		case strategies.Random:
			nextMove = applyRandomStrategy(validMoves)
		case strategies.RandomHVFirstDiagRandom:
			nextMove = applyRandomHVFirstDiagRandom(validMoves)
		}

		return &nextMove, models.StatusErr{}
	}

	return nil, models.StatusErr{GameStatus: models.NoMoreMoves, Message: "No more free moves"}

}

func getRemainingMoves(moves map[models.Movement]bool, movesList []models.Movement) ([]models.Movement, int) {
	var validMoves []models.Movement
	for _, m := range movesList {
		if !moves[m] {
			validMoves = append(validMoves, m)
		}
	}
	return validMoves, len(validMoves)
}

func applyRandomStrategy(validMoves []models.Movement) models.Movement {
	return validMoves[rand.Intn(len(validMoves))]
}

func applyRandomHVFirstDiagRandom(validMoves []models.Movement) models.Movement {
	var crossMoves = []models.Movement{models.Up, models.Down, models.Right, models.Left}
	var diagMoves = []models.Movement{models.UpLeft, models.UpRight, models.DownLeft, models.DownRight}

	validCrossMoves := make([]models.Movement, 0, 4)
	for _, dir := range crossMoves {
		if containsMove(validMoves, dir) {
			validCrossMoves = append(validCrossMoves, dir)
		}
	}

	validCrossMovesSize := len(validCrossMoves)
	if validCrossMovesSize != 0 {
		return validCrossMoves[rand.Intn(validCrossMovesSize)]
	}

	//no more cross moves, go to diagonal ones

	validDiagMoves := make([]models.Movement, 0, 4)
	for _, dir := range diagMoves {
		if containsMove(validMoves, dir) {
			validDiagMoves = append(validDiagMoves, dir)
		}
	}

	validDiagMovesSize := len(validDiagMoves)

	// no need to check for the size, 'cause it is a pre-condition that at least we have 1 valid move
	return validDiagMoves[rand.Intn(validDiagMovesSize)]
}

func containsMove(validMoves []models.Movement, dir models.Movement) bool {
	for _, item := range validMoves {
		if item == dir {
			return true
		}
	}
	return false
}
