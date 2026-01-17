package sidcon

import "testing"

func TestParseCube(t *testing.T) {
	var inputs []string = []string{"w", "g", "b", "K", "B", "Y", "UT", "sh"}
	var expected []Cube = []Cube{White, Green, Brown, Black, Blue, Yellow, Ultra, Ship}
	for i := range inputs {
		cube, err := ParseCube(inputs[i])
		if cube != expected[i] || err != nil {
			t.Errorf(`ParseCube(%q) = %q, %v, want %v, nil`, inputs[i], cube, err, expected[i])
		}
	}
}

func TestParseCubeError(t *testing.T) {
	var inputs []string = []string{"", " ", "gg", "W", " b ", "+K", "-B", "1Y", "U T", "hs"}
	for _, input := range inputs {
		cube, err := ParseCube(input)
		if cube != -1 || err == nil {
			t.Errorf(`ParseCube(%q) = %q, %v, want -1, error`, input, cube, err)
		}
	}
}
