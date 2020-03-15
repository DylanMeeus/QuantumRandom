package main

import (
	"fmt"
	"github.com/DylanMeeus/QuantumRandom/pkg"
)

func main() {
	n8, err := pkg.NextUint8()
	if err != nil {
		panic(err)
	}
	fmt.Printf("random uint8 value: %v\n", n8)

	i, err := pkg.NextInt()
	if err != nil {
		panic(err)
	}
	fmt.Printf("random int value: %v\n", i)
}
