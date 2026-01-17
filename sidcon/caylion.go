package sidcon

type Caylion struct {
	name          string
	cubes         Cubes
	converters    []Converter
	colonies      []Colony
	colonySupport int
	bidTiebreaker int
}

func (c Caylion) Init() {
	c.name = "Caylion"
	c.cubes = Cubes{Black: 2, Green: 5, White: 4, Brown: 2, Ship: 1}
	c.converters = []Converter{
		Converter{
			"Planetary Ecological Dominance",
			Cubes{},
			Cubes{Ship: 1, Green: 3, White: 1},
		},
		Converter{
			"Oceanic Processing",
			Cubes{Green: 2},
			Cubes{Blue: 1, Brown: 1},
		},
		Converter{
			"Distributed Telenet",
			Cubes{Black: 1},
			Cubes{Brown: 1, White: 1},
		},
		Converter{
			"Lunar Mining Consortium",
			Cubes{Blue: 2},
			Cubes{Ship: 2, Black: 1, White: 1},
		},
	}
	var colony Colony = GenerateColony()
	colony.hasSeedlingToken = true
	c.colonies = []Colony{colony}
}

func (c Caylion) Name() string {
	return c.name
}

func (c Caylion) Cubes() Cubes {
	return c.cubes
}

func (c Caylion) Converters() []Converter {
	return c.converters
}

func (c Caylion) Colonies() []Colony {
	return c.colonies
}

func (c Caylion) ColonySupport() int {
	return c.colonySupport
}

func (c Caylion) BidTiebreaker() int {
	return c.bidTiebreaker
}
