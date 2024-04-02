package gameof100

type HorizontalMovement interface {
	MoveLeft() *GameCell
	MoveRight() *GameCell
}

type VerticalMovement interface {
	MoveTop() *GameCell
	MoveBottom() *GameCell
}

type DiagonalMovement interface {
	MoveTopLeft() *GameCell
	MoveTopRight() *GameCell
	MoveBottomLeft() *GameCell
	MoveBottomRight() *GameCell
}


type GameCell struct {
	val uint8
	xPos uint8
	yPos uint8
}

type Grid struct {
	lines [][]GameCell
}
