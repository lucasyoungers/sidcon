package sidcon

type Faderan struct {
	props FactionProperties
}

func (f Faderan) Init() {
	f.props.name = "Faderan"
	f.props.cubes = Cubes{Ultra: 1, Yellow: 1, Blue: 1, White: 3, Brown: 1, Ship: 2}
	f.props.converters = []Converter{
		{
			"Archaic Automated Support Network",
			Cubes{},
			Cubes{Ship: 1, White: 1, Green: 1, Brown: 1},
		},
		{
			"Relic Repurposing",
			Cubes{Brown: 2},
			Cubes{Ship: 2, White: 1},
		},
		{
			"Non-Indigenous Fauna",
			Cubes{Green: 3},
			Cubes{Blue: 1, Yellow: 1, Brown: 1},
		},
		{
			"Ritual Government",
			Cubes{White: 4},
			Cubes{Ultra: 1, Black: 1, Green: 1},
		},
	}
}

func (f Faderan) Properties() FactionProperties {
	return f.props
}
