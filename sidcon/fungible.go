package sidcon

type Fungible interface {
	AddTo(Faction)
	RemoveFrom(Faction) error
}
