package sidcon

import (
	"fmt"
	"strings"
)

type Converter struct {
	name   string
	input  Cubes
	output Cubes
}

func (c Converter) String() string {
	return fmt.Sprintf("%s: %s -> %s", c.name, c.input, c.output)
}

func (c Converter) Input() Cubes {
	return c.input
}

func (c Converter) Output() Cubes {
	return c.output
}

func (c Converter) AddTo(f Faction) {
	f.Properties().AddConverter(c)
}

func (c Converter) RemoveFrom(f Faction) error {
	return f.Properties().RemoveConverter(c)
}

func ParseConverter(s string) (Converter, error) {
	var nameSplit []string = strings.Split(s, ":")
	if len(nameSplit) < 2 {
		return Converter{}, fmt.Errorf("converter malformed")
	}
	var name string = nameSplit[0]

	var converterSplit []string = strings.Split(nameSplit[1], "->")
	if len(converterSplit) < 2 {
		return Converter{}, fmt.Errorf("converter body malformed")
	}

	input, err := ParseCubes(converterSplit[0])
	if err != nil {
		return Converter{}, fmt.Errorf("converter input malformed")
	}

	output, err := ParseCubes(converterSplit[1])
	if err != nil {
		return Converter{}, fmt.Errorf("converter output malformed")
	}

	return Converter{name, input, output}, nil
}
