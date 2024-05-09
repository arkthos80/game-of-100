package models

type Cell struct {
	Val  string
	xPos int
	yPos int
}

func (c Cell) GetPosition() (int, int){
	return c.xPos,c.yPos
}

func MakeCellPtr(val string, xPos, yPos int) *Cell {
	return &Cell{Val: val, xPos: xPos, yPos: yPos}
}