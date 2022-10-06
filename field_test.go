package life

import "testing"

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
