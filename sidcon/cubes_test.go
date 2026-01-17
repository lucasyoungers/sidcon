package sidcon

import "testing"

func TestCubesString(t *testing.T) {
	var inputs = []Cubes{
		{Ultra: 1},
		{Yellow: 3},
		{Brown: 2, Green: 3, White: 4},
	}
	var expected = []string{
		"1UT",
		"3Y",
		"4w, 3g, 2b",
	}
	for i := range inputs {
		output := inputs[i].String()
		if inputs[i].String() != expected[i] {
			t.Errorf("(%q).String() = %q, want %q", inputs[i], output, expected[i])
		}
	}
}

func TestCubesAddCubes(t *testing.T) {
	var inputLeft Cubes = Cubes{White: 4, Green: 3, Brown: 2}
	var inputRight Cubes = Cubes{White: 1, Green: 6, Ultra: 1, Ship: 4}
	var output Cubes = Cubes{White: 4, Green: 3, Brown: 2}
	var expected Cubes = Cubes{White: 5, Green: 9, Brown: 2, Ultra: 1, Ship: 4}
	output.AddCubes(inputRight)
	if output.String() != expected.String() {
		t.Errorf("(%s).AddCubes(%s) = %q, want %q", inputLeft, inputRight, output, expected)
	}
}

func TestCubesRemoveCubes(t *testing.T) {
	var inputLeft Cubes = Cubes{White: 4, Green: 3, Brown: 2}
	var inputRight Cubes = Cubes{White: 1, Green: 3}
	var output Cubes = Cubes{White: 4, Green: 3, Brown: 2}
	var expected Cubes = Cubes{White: 3, Brown: 2}
	err := output.RemoveCubes(inputRight)
	if output.String() != expected.String() || err != nil {
		t.Errorf("(%s).RemoveCubes(%s) = %q, want %q", inputLeft, inputRight, output, expected)
	}
}

func TestCubesRemoveCubesError(t *testing.T) {
	var inputLeft Cubes = Cubes{White: 4, Green: 3, Brown: 2}
	var inputRight Cubes = Cubes{White: 1, Green: 3, Brown: 5}
	var output Cubes = Cubes{White: 4, Green: 3, Brown: 2}
	err := output.RemoveCubes(inputRight)
	if err == nil {
		t.Errorf("(%s).RemoveCubes(%s) = %q, expected error", inputLeft, inputRight, output)
	}
}

func TestParseCubes(t *testing.T) {
	var inputs []string = []string{
		"2b, 3g, 4w",
		"1UT, 100w, 80000sh",
		"2b3g4w",
		"@2b-3g#4w?",
	}
	var expected []Cubes = []Cubes{
		{White: 4, Green: 3, Brown: 2},
		{White: 100, Ultra: 1, Ship: 80000},
		{White: 4, Green: 3, Brown: 2},
		{White: 4, Green: 3, Brown: 2},
	}
	for i := range inputs {
		output, err := ParseCubes(inputs[i])
		status := true
		for cube, quantity := range expected[i] {
			if output[cube] != quantity {
				status = false
			}
		}
		if !status || len(output) != len(expected[i]) || err != nil {
			t.Errorf("ParseCubes(%q) = %q, %v, want %q, nil", inputs[i], output, err, expected[i])
		}
	}
}
