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
	fmt.Printf("random number: %v\n", n8)
}
