package sidcon

import "fmt"

type Cube int

const (
	White Cube = iota
	Green
	Brown
	Black
	Blue
	Yellow
	Ultra
	Ship
)

var color = map[Cube]string{
	White:  "w",
	Green:  "g",
	Brown:  "b",
	Black:  "K",
	Blue:   "B",
	Yellow: "Y",
	Ultra:  "UT",
	Ship:   "sh",
}

var value = map[Cube]float32{
	White:  1,
	Green:  1,
	Brown:  1,
	Black:  1.5,
	Blue:   1.5,
	Yellow: 1.5,
	Ultra:  3,
	Ship:   1,
}

func (c Cube) String() string {
	return color[c]
}

func (c Cube) Value() float32 {
	return value[c]
}

func ParseCube(s string) (Cube, error) {
	for cube, color := range color {
		if color == s {
			return cube, nil
		}
	}
	return -1, fmt.Errorf("%s is not a valid color", s)
}
