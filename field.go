package life

type Field [][]bool

func NewField(width int, height int) Field {
	field := Field{}
	for j := 0; j < height; j++ {
		row := []bool{}
		for i := 0; i < width; i++ {
			row = append(row, false)
		}

		field = append(field, row)
	}

	return field
}

func (field Field) Width() int {
	return len(field[0])
}

func (field Field) Height() int {
	return len(field)
}

func (field Field) Cell(column int, row int) bool {
	column = (column + field.Width()) % field.Width()
	row = (row + field.Height()) % field.Height()
	return field[row][column]
}

func (field Field) SetCell(column int, row int, cell bool) {
	field[row][column] = cell
}

func (field Field) NeighborCount(column int, row int) int {
	var count int
	for j := row - 1; j <= row+1; j++ {
		for i := column - 1; i <= column+1; i++ {
			if field.Cell(i, j) {
				count++
			}
		}
	}

	if field.Cell(column, row) {
		count--
	}

	return count
}
