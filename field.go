package life

type Field [][]bool

func (field Field) Width() int {
	return len(field[0])
}
