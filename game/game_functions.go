package game

import (
	"game-of-100/models"
	"game-of-100/strategies"
	"math"
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

	nextMovePtr, err = getNextMoveByStrategy(gameTable.GridSize(), currCellPtr, allMovesMap, strategy)
	tries := 1

	for err.GameStatus != models.NoMoreMoves {
		currCellPtr, moveErr := gameTable.Move(currCellPtr, *nextMovePtr)
		if moveErr == nil {
			//found a valid move, return it
			return currCellPtr, *nextMovePtr, tries, models.StatusErr{}
		} else {
			allMovesMap[*nextMovePtr] = true
			tries += 1
		}

		if currCellPtr == nil && err.GameStatus != models.NoMoreMoves {
			// an error occurred, but still not an end
			// restoring original cell
			currCellPtr = fromCellPtr
		}

		nextMovePtr, err = getNextMoveByStrategy(gameTable.GridSize(), currCellPtr, allMovesMap, strategy)
	}

	return fromCellPtr, -1, tries, models.StatusErr{GameStatus: models.NoMoreMoves, Message: "No more free moves"}
}

func getNextMoveByStrategy(gridSize int8, currCellPtr *models.Cell, moves map[models.Movement]bool, strategy strategies.Strategy) (*models.Movement, models.StatusErr) {
	validMoves, sizeValidMoves := getValidMoves(moves, models.AllMoves)

	if sizeValidMoves != 0 {
		var nextMove models.Movement
		switch strategy {
		case strategies.Random:
			nextMove = applyRandom(validMoves)
		case strategies.CrossFirst:
			nextMove = applyCrossFirst(validMoves)
		case strategies.DiagonalFirst:
			nextMove = applyDiagonalFirst(validMoves)
		case strategies.CloseToBorder:
			nextMove = applyCloseToBorder(gridSize, currCellPtr, validMoves)
		default:
			return nil, models.StatusErr{GameStatus: models.NoMoreMoves, Message: "Can't proceed. Chosen strategy is not implemented"}
		}

		return &nextMove, models.StatusErr{}
	}

	return nil, models.StatusErr{GameStatus: models.NoMoreMoves, Message: "No more free moves"}

}

func getValidMoves(moves map[models.Movement]bool, movesList []models.Movement) ([]models.Movement, int) {
	var validMoves []models.Movement
	for _, m := range movesList {
		if !moves[m] {
			validMoves = append(validMoves, m)
		}
	}
	return validMoves, len(validMoves)
}

func applyRandom(validMoves []models.Movement) models.Movement {
	return validMoves[rand.Intn(len(validMoves))]
}

func applyCrossFirst(validMoves []models.Movement) models.Movement {

	validCrossMoves, validCrossMovesSize := remainingMoves(validMoves, models.CrossMoves)

	if validCrossMovesSize != 0 {
		return validCrossMoves[rand.Intn(validCrossMovesSize)]
	}

	//no more cross moves, go to diagonal ones

	validDiagMoves, validDiagMovesSize := remainingMoves(validMoves, models.DiagonalMoves)

	// no need to check for the size, 'cause it is a pre-condition that at least we have 1 valid move
	return validDiagMoves[rand.Intn(validDiagMovesSize)]
}

func applyDiagonalFirst(validMoves []models.Movement) models.Movement {
	validDiagMoves, validDiagMovesSize := remainingMoves(validMoves, models.DiagonalMoves)

	if validDiagMovesSize != 0 {
		return validDiagMoves[rand.Intn(validDiagMovesSize)]
	}

	//no more diagonal moves, go to cross ones

	validCrossMoves, validCrossMovesSize := remainingMoves(validMoves, models.CrossMoves)

	// no need to check for the size, 'cause it is a pre-condition that at least we have 1 valid move
	return validCrossMoves[rand.Intn(validCrossMovesSize)]

}

func applyCloseToBorder(gridSize int8, currCellPtr *models.Cell, validMoves []models.Movement) models.Movement {
	x, y := currCellPtr.GetPosition()

	closerX, closerY := x, y
	closerMoveIdx := -1

	for i, dir := range validMoves {
		newX, newY := models.GetNextPositionByMove(dir, x, y)
		if !models.IsOutOfBound(newX, newY, gridSize) && isCloserToBorder(gridSize, newX, newY, x, y) {

			if isCloserToBorder(gridSize, newX, newY, closerX, closerY) { //taking the min closer between all eligible moves
				closerX = newX
				closerY = newY
				closerMoveIdx = i
			}
		}
	}

	if closerMoveIdx == -1 {
		//no closer move found, take a random valid one
		closerMoveIdx = rand.Intn(len(validMoves))
	}

	return validMoves[closerMoveIdx]
}

func isCloserToBorder(gridSize int8, newX, newY, oldX, oldY int8) bool {
	// Calculate distances to borders
	newDistX := minimum(newX, gridSize-newX) // Distance to left or right border
	newDistY := minimum(newY, gridSize-newY) // Distance to top or bottom border
	oldDistX := minimum(oldX, gridSize-oldX)
	oldDistY := minimum(oldY, gridSize-oldY)

	// Compare distances
	return newDistX < oldDistX || newDistY < oldDistY
}

func minimum(a, b int8) float64 {
	return math.Min(float64(a), float64(b))
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
