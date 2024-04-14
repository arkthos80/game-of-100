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
		case strategies.CrossFirstRandom:
			nextMove = applyCrossFirstRandom(validMoves)
		case strategies.DiagonalFirstRandom:
			nextMove = applyDiagonalFirstRandom(validMoves)
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

func applyCrossFirstRandom(validMoves []models.Movement) models.Movement {

	validCrossMoves, validCrossMovesSize := remainingMoves(validMoves, models.CrossMoves)

	if validCrossMovesSize != 0 {
		return validCrossMoves[rand.Intn(validCrossMovesSize)]
	}

	//no more cross moves, go to diagonal ones

	validDiagMoves, validDiagMovesSize := remainingMoves(validMoves, models.DiagonalMoves)

	// no need to check for the size, 'cause it is a pre-condition that at least we have 1 valid move
	return validDiagMoves[rand.Intn(validDiagMovesSize)]
}

func applyDiagonalFirstRandom(validMoves []models.Movement) models.Movement {
	validDiagMoves, validDiagMovesSize := remainingMoves(validMoves, models.DiagonalMoves)

	if validDiagMovesSize != 0 {
		return validDiagMoves[rand.Intn(validDiagMovesSize)]
	}

	//no more diagonal moves, go to cross ones

	validCrossMoves, validCrossMovesSize := remainingMoves(validMoves, models.CrossMoves)

	// no need to check for the size, 'cause it is a pre-condition that at least we have 1 valid move
	return validCrossMoves[rand.Intn(validCrossMovesSize)]

}

func remainingMoves(moves []models.Movement, movesSet []models.Movement) ([]models.Movement, int) {
	validMoves := make([]models.Movement, 0, 4)

	for _, dir := range movesSet {
		if containsMove(moves, dir) {
			validMoves = append(validMoves, dir)
		}
	}

	validMovesSize := len(validMoves)
	return validMoves, validMovesSize
}

func containsMove(validMoves []models.Movement, dir models.Movement) bool {
	for _, item := range validMoves {
		if item == dir {
			return true
		}
	}
	return false
}
