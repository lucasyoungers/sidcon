package sidcon

type Kjas struct {
	props FactionProperties
}

func (f Kjas) Init() {
	f.props.name = "Kjas"
	f.props.cubes = Cubes{Blue: 1, Yellow: 2, Black: 2, Ship: 4, Green: 1, White: 1, Brown: 2}
	f.props.converters = []Converter{
		{
			"Underground Cities",
			Cubes{},
			Cubes{Ship: 3, Brown: 2},
		},
		{
			"Perfume Pheromone Loyalty",
			Cubes{Brown: 4},
			Cubes{Yellow: 1, Black: 1, Blue: 1, White: 1},
		},
		{
			"Starstrider Jumpcores",
			Cubes{Blue: 2},
			Cubes{Ship: 4, Green: 1},
		},
	}
	f.props.colonies = []Colony{GenerateColony(), GenerateColony()}
	f.props.colonySupport = 6
	f.props.bidTiebreaker = 6
}

func (f Kjas) Properties() FactionProperties {
	return f.props
}

func (f Kjas) SetProperties(props FactionProperties) {
	f.props = props
}
