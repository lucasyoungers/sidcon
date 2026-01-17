package main

import (
	"fmt"

	sc "github.com/lucasyoungers/sidcon/sidcon"
)

func main() {
	var c sc.Cubes = sc.Cubes{
		sc.White: 4,
		sc.Green: 3,
		sc.Brown: 2,
	}
	fmt.Println(c)
}
