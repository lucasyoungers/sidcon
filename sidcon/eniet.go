package sidcon

type EniEt struct {
	props FactionProperties
}

func (f EniEt) Init() {
	f.props.name = "Eni Et"
	f.props.cubes = Cubes{Blue: 1, Brown: 2, White: 2, Green: 1, Ship: 1}
	f.props.converters = []Converter{
		{
			"Deep Hunting",
			Cubes{},
			Cubes{Ship: 1, Blue: 1},
		},
		{
			"Choral Song",
			Cubes{White: 5},
			Cubes{Ship: 2, Ultra: 1, Black: 1, Green: 1},
		},
		{
			"Deep Hunting",
			Cubes{Blue: 2},
			Cubes{White: 2, Brown: 1, Green: 1},
		},
		// {
		// 	"Universal Applied Metaethics",
		// 	Cubes{},
		// 	Cubes{},
		// },
	}
	f.props.colonies = []Colony{GenerateColony()}
	f.props.colonySupport = 3
	f.props.bidTiebreaker = 3
}
