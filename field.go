package life

type Field [][]bool

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
