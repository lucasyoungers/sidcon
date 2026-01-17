package sidcon

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
