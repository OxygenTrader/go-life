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

func TestField_NeighbourCount(test *testing.T) {
	field := Field{
		{true, false, false},
		{false, true, true},
		{true, false, false},
	}
	actualCount := field.NeighborCount(0, 1)

	expectedCount := 4
	if actualCount != expectedCount {
		test.Fail()
	}
}

func TestField_NeighbourCount_forAliveCell(test *testing.T) {
	field := Field{
		{true, false, false},
		{true, true, true},
		{true, false, false},
	}
	actualCount := field.NeighborCount(0, 1)

	expectedCount := 4
	if actualCount != expectedCount {
		test.Fail()
	}
}

func TestField_NextCell_withBirth(test *testing.T) {
	field := Field{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, false, true, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
	}
	actualCell := field.NextCell(1, 2)

	expectedCell := true
	if actualCell != expectedCell {
		test.Fail()
	}
}

func TestField_NextCell_withSurvival(test *testing.T) {
	field := Field{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, false, true, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
	}
	actualCell := field.NextCell(2, 3)

	expectedCell := true
	if actualCell != expectedCell {
		test.Fail()
	}
}

func TestField_NextCell_withDeath(test *testing.T) {
	field := Field{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, false, true, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
	}
	actualCell := field.NextCell(1, 3)

	expectedCell := false
	if actualCell != expectedCell {
		test.Fail()
	}
}
