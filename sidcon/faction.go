package sidcon

import "fmt"

type Faction interface {
	Init()
	Properties() FactionProperties
	SetProperties(FactionProperties)
}

type FactionProperties struct {
	name          string
	cubes         Cubes
	converters    []Converter
	colonies      []Colony
	colonySupport int
	bidTiebreaker int
}

func (p FactionProperties) AddCubes(cubes Cubes) {
	p.cubes.AddCubes(cubes)
}

func (p FactionProperties) RemoveCubes(cubes Cubes) error {
	return p.cubes.RemoveCubes(cubes)
}

func (p FactionProperties) AddConverter(converter Converter) {
	p.converters = append(p.converters, converter)
}

func (p FactionProperties) RemoveConverter(converter Converter) error {
	var removeIndex int = -1
	for i, c := range p.converters {
		if c.name == converter.name {
			removeIndex = i
			break
		}
	}
	if removeIndex == -1 {
		return fmt.Errorf("converter %q not found", converter.name)
	}
	p.converters = append(p.converters[:removeIndex], p.converters[removeIndex+1:]...)
	return nil
}

func (p FactionProperties) AddColony(colony Colony) {
	p.colonies = append(p.colonies, colony)
}

func (p FactionProperties) RemoveColony(colony Colony) error {
	var removeIndex int = -1
	for i, c := range p.colonies {
		if c.name == colony.name {
			removeIndex = i
			break
		}
	}
	if removeIndex == -1 {
		return fmt.Errorf("colony %q not found", colony.name)
	}
	p.colonies = append(p.colonies[:removeIndex], p.colonies[removeIndex+1:]...)
	return nil
}

func (p FactionProperties) RunFactory(factory Factory) error {
	err := p.cubes.RemoveCubes(factory.Input())
	if err != nil {
		return err
	}
	p.cubes.AddCubes(factory.Output())
	return nil
}

func (p FactionProperties) RunFactories(factories []Factory) error {
	var inputs Cubes
	var outputs Cubes
	for _, factory := range factories {
		inputs.AddCubes(factory.Input())
		outputs.AddCubes(factory.Output())
	}
	err := p.cubes.RemoveCubes(inputs)
	if err != nil {
		return err
	}
	p.cubes.AddCubes(outputs)
	return nil
}
