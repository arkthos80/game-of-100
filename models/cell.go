package models

type Cell struct {
	Val  string
	xPos int8
	yPos int8
}

func MakeCellPtr(val string, xPos, yPos int8) *Cell {
	return &Cell{Val: val, xPos: xPos, yPos: yPos}
}