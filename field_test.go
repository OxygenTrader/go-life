package life

import (
	"reflect"
	"testing"
)

func TestNewField(test *testing.T) {
	field := NewField(3, 2)

	expectedField := Field{
		{false, false, false},
		{false, false, false},
	}
	if !reflect.DeepEqual(field, expectedField) {
		test.Fail()
	}
}

func TestField_Width(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, false},
	}
	actualWidth := field.Width()

	expectedWidth := 3
	if actualWidth != expectedWidth {
		test.Fail()
	}
}

func TestField_Height(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, false},
	}
	actualHeight := field.Height()

	expectedHeight := 2
	if actualHeight != expectedHeight {
		test.Fail()
	}
}

func TestField_Cell(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, true},
	}
	actualCell := field.Cell(2, 1)

	expectedCell := true
	if actualCell != expectedCell {
		test.Fail()
	}
}

func TestField_Cell_withCoordinatesBeyondMinimum(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, true, false},
	}
	actualCell := field.Cell(-2, -1)

	expectedCell := true
	if actualCell != expectedCell {
		test.Fail()
	}
}

func TestField_Cell_withCoordinatesBeyondMaximum(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, true, false},
	}
	actualCell := field.Cell(4, 3)

	expectedCell := true
	if actualCell != expectedCell {
		test.Fail()
	}
}

func TestField_SetCell(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, false},
	}
	field.SetCell(2, 1, true)

	expectedField := Field{
		{false, false, false},
		{false, false, true},
	}
	if !reflect.DeepEqual(field, expectedField) {
		test.Fail()
	}
}
