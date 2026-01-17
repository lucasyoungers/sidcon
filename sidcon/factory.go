package sidcon

type Factory interface {
	Input() Cubes
	Output() Cubes
}
