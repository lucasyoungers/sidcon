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
