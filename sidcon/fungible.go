package sidcon

type Fungible interface {
	RemoveFrom(Faction)
	AddTo(Faction)
}
