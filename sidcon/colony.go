package sidcon

import (
	"fmt"
	"math/rand/v2"
)

type Climate int

const (
	Jungle Climate = iota
	Desert
	Ocean
	Ice
)

var climate = map[Climate]string{
	Jungle: "Jungle",
	Desert: "Desert",
	Ocean:  "Ocean",
	Ice:    "Ice",
}

func (c Climate) String() string {
	return climate[c]
}

type Colony struct {
	name             string
	climate          Climate
	input            Cubes
	output           Cubes
	hasSeedlingToken bool
	isKtPlanet       bool
}

func (c Colony) String() string {
	var prefix string

	if c.hasSeedlingToken {
		prefix += "Seedling "
	}

	if c.isKtPlanet {
		prefix += "Kt "
	}

	return prefix + fmt.Sprintf("%s Colony: %s -> %s", c.climate.String(), c.input.String(), c.output.String())
}

func (c Colony) Input() Cubes {
	return c.input
}

func (c Colony) Output() Cubes {
	return c.output
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func generateName() string {
	b := make([]rune, 10)
	for i := range 10 {
		if rand.IntN(5) == 0 {
			b[i] = ' '
		} else {
			b[i] = letters[rand.IntN(len(letters))]
		}
	}
	return string(b)
}

func GenerateColony() Colony {
	var climate Climate = Climate(rand.IntN(len(climate)))
	var cube Cube = Cube(rand.IntN(len(color)))
	return Colony{generateName(), climate, Cubes{}, Cubes{cube: 1}, false, false}
}
