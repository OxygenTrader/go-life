package life

import (
	"errors"
	"fmt"
)

type Field [][]bool

func NewField(width int, height int) Field {
	var field Field
	for j := 0; j < height; j++ {
		var row []bool
		for i := 0; i < width; i++ {
			row = append(row, false)
		}

		field = append(field, row)
	}

	return field
}

func ParseField(text string) (Field, error) {
	var field Field
	var row []bool
	var isInsideComment bool
	for _, character := range text {
		if isInsideComment {
			if character == '\n' {
				isInsideComment = false
			}

			continue
		}

		switch character {
		case '.':
			row = append(row, false)
		case '0':
			row = append(row, true)
		case '\n':
			hasSomeRows := field.Height() != 0
			if hasSomeRows && len(row) != field.Width() {
				return nil, errors.New("inconsistent row length")
			}

			field = append(field, row)
			row = []bool{}
		case '!':
			isInsideComment = true
		default:
			return nil, fmt.Errorf("unknown character %q", character)
		}
	}

	return field, nil
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
	for neighborRow := row - 1; neighborRow <= row+1; neighborRow++ {
		for neighborColumn := column - 1; neighborColumn <= column+1; neighborColumn++ {
			if field.Cell(neighborColumn, neighborRow) {
				count++
			}
		}
	}

	if field.Cell(column, row) {
		count--
	}

	return count
}

func (field Field) NextCell(column int, row int) bool {
	cell := field.Cell(column, row)
	neighborCount := field.NeighborCount(column, row)

	willBeBorn := !cell && neighborCount == 3
	WillSurvive := cell && (neighborCount == 2 || neighborCount == 3)

	return willBeBorn || WillSurvive
}

func (field Field) NextField() Field {
	nextField := NewField(field.Width(), field.Height())
	for row := 0; row < field.Height(); row++ {
		for column := 0; column < field.Width(); column++ {
			nextCell := field.NextCell(column, row)
			nextField.SetCell(column, row, nextCell)
		}
	}

	return nextField
}

func (field Field) String() string {
	var result string
	for row := 0; row < field.Height(); row++ {
		for column := 0; column < field.Width(); column++ {
			if field.Cell(column, row) {
				result += "0"
			} else {
				result += "."
			}
		}

		result += "\n"
	}

	return result
}
