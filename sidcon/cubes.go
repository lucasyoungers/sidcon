package sidcon

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Cubes map[Cube]int

func (c Cubes) String() string {
	var output []string = make([]string, len(c))
	for i, cube := range slices.Sorted(maps.Keys(c)) {
		output[i] = fmt.Sprintf("%d%s", c[cube], cube)
	}
	return strings.Join(output, ", ")
}

func (c Cubes) Value() float32 {
	var output float32
	for cube, quantity := range c {
		output += (cube.Value() * float32(quantity))
	}
	return output
}

func (c Cubes) Add(cube Cube, quantity int) {
	c[cube] += quantity
}

func (c Cubes) Remove(cube Cube, quantity int) {
	c[cube] -= quantity
}

func (c Cubes) AddCubes(other Cubes) {
	for cube, quantity := range other {
		c.Add(cube, quantity)
	}
}

func (c Cubes) RemoveCubes(other Cubes) {
	for cube, quantity := range other {
		c.Remove(cube, quantity)
	}
}

func (c Cubes) Contains(cube Cube, quantity int) bool {
	return c[cube] >= quantity
}

func (c Cubes) ContainsCubes(other Cubes) bool {
	for cube, quantity := range other {
		if !c.Contains(cube, quantity) {
			return false
		}
	}
	return true
}

func ParseCubes(s string) (Cubes, error) {
	var reMain *regexp.Regexp = regexp.MustCompile(`\d+[A-Za-z]+`)
	var reQuantity *regexp.Regexp = regexp.MustCompile(`\d+`)
	var reColor *regexp.Regexp = regexp.MustCompile(`[A-Za-z]+`)
	var matches []string = reMain.FindAllString(s, -1)
	var output Cubes = make(Cubes)
	for _, match := range matches {
		strQuantity := reQuantity.FindString(match)
		strColor := reColor.FindString(match)
		cube, err := ParseCube(strColor)
		if err != nil {
			return nil, err
		}
		quantity, err := strconv.Atoi(strQuantity)
		if err != nil {
			return nil, fmt.Errorf("malformed cube string")
		}
		output[cube] = quantity
	}
	return output, nil
}
