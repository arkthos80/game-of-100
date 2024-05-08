package models

type Cell struct {
	Val  string
	xPos int8
	yPos int8
}

func (c Cell) GetPosition() (int8, int8){
	return c.xPos,c.yPos
}

func MakeCellPtr(val string, xPos, yPos int8) *Cell {
	return &Cell{Val: val, xPos: xPos, yPos: yPos}
}