package game

import (
	"game-of-100/models"
)

func InitGameTable(n int, defaultVal string) models.Table {
	var lines [][]*models.Cell

	for r := 0; r < n; r++ {
		var line []*models.Cell
		for i := 0; i < n; i++ {
			line = append(line, models.MakeCellPtr(defaultVal, int8(i), int8(r)))
		}
		lines = append(lines, line)
	}

	return models.MakeTable(lines)
}
