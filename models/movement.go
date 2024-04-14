package models

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

var AllMoves = []Movement{Up, Down, Left, Right, DownRight, DownLeft, UpLeft, UpRight}
var CrossMoves = []Movement{Up, Down, Right, Left}
var DiagonalMoves = []Movement{UpLeft, UpRight, DownLeft, DownRight}

var AllMovesNames = [...]string{
	"Up",
	"Down",
	"Left",
	"Right",
	"UpLeft",
	"UpRight",
	"DownLeft",
	"DownRight",
}