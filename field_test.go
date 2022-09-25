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
