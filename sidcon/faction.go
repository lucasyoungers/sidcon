package sidcon

type Faction interface {
	Init()
	Name() string
	Cubes() Cubes
	Converters() []Converter
	Colonies() []Colony
	ColonySupport() int
	BidTiebreaker() int
}
