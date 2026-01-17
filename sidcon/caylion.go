package sidcon

type Caylion struct {
	props FactionProperties
}

func (f Caylion) Init() {
	f.props.name = "Caylion"
	f.props.cubes = Cubes{Black: 2, Green: 5, White: 4, Brown: 2, Ship: 1}
	f.props.converters = []Converter{
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
	f.props.colonies = []Colony{colony}
	f.props.colonySupport = 3
	f.props.bidTiebreaker = 1
}

func (f Caylion) Properties() FactionProperties {
	return f.props
}

func (f Caylion) SetProperties(props FactionProperties) {
	f.props = props
}
