package sidcon

type Caylion struct {
	name          string
	cubes         Cubes
	converters    []Converter
	colonies      []Colony
	colonySupport int
	bidTiebreaker int
}

func (f Caylion) Init() {
	f.name = "Caylion"
	f.cubes = Cubes{Black: 2, Green: 5, White: 4, Brown: 2, Ship: 1}
	f.converters = []Converter{
		{
			"Planetary Ecological Dominance",
			Cubes{},
			Cubes{Ship: 1, Green: 3, White: 1},
		},
		{
			"Oceanic Processing",
			Cubes{Green: 2},
			Cubes{Blue: 1, Brown: 1},
		},
		{
			"Distributed Telenet",
			Cubes{Black: 1},
			Cubes{Brown: 1, White: 1},
		},
		{
			"Lunar Mining Consortium",
			Cubes{Blue: 2},
			Cubes{Ship: 2, Black: 1, White: 1},
		},
	}
	var colony Colony = GenerateColony()
	colony.hasSeedlingToken = true
	f.colonies = []Colony{colony}
}

func (f Caylion) Name() string {
	return f.name
}

func (f Caylion) Cubes() Cubes {
	return f.cubes
}

func (f Caylion) Converters() []Converter {
	return f.converters
}

func (f Caylion) Colonies() []Colony {
	return f.colonies
}

func (f Caylion) ColonySupport() int {
	return f.colonySupport
}

func (f Caylion) BidTiebreaker() int {
	return f.bidTiebreaker
}
