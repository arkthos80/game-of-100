package models

type Status int

const (
	InvalidPosition Status = iota + 1
	InvalidCellValue
	CellAlreadyUsed
	NoMoreMoves
)
