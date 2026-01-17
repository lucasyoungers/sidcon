package sidcon

import "testing"

func TestParseConverter(t *testing.T) {
	var input string = "Nanotechnology: 3b -> 1UT, 1w"
	var expected Converter = Converter{
		"Nanotechnology",
		Cubes{Brown: 3},
		Cubes{Ultra: 1, White: 1},
	}
	output, err := ParseConverter(input)
	if output.String() != expected.String() || err != nil {
		t.Errorf("ParseConverter(%q) = %q, %v, want %q, nil", input, output, err, expected)
	}
}
