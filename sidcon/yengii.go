package sidcon

type Yengii struct {
	props FactionProperties
}

func (f Yengii) Init() {
	f.props.name = "Yengii"
	f.props.cubes = Cubes{Ultra: 1, Blue: 1, Yellow: 1, Black: 1, Green: 1, White: 1, Brown: 1}
	f.props.converters = []Converter{
		{
			"Social Ecological Mutualism",
			Cubes{},
			Cubes{Ship: 2, Blue: 1, Brown: 1},
		},
		{
			"Choice of Paths",
			Cubes{White: 2, Black: 1},
			Cubes{Ultra: 1, Green: 2},
		},
		{
			"Empathy of Worlds",
			Cubes{Brown: 2, Blue: 1},
			Cubes{Black: 2, White: 2},
		},
		{
			"Work of Mules",
			Cubes{Green: 2, Yellow: 1},
			Cubes{Ship: 2, Blue: 1, Brown: 2},
		},
		{
			"Social Ecological Mutualism",
			Cubes{Ultra: 1},
			Cubes{Ship: 2, Yellow: 1, Green: 1},
		},
	}
	f.props.colonies = []Colony{GenerateColony()}
	f.props.colonySupport = 3
	f.props.bidTiebreaker = 5
}

func (f Yengii) Properties() FactionProperties {
	return f.props
}
