package main

import (
	"fmt"
	"github.com/DylanMeeus/QuantumRandom/pkg"
)

func main() {
	fmt.Println("vim-go")
	n8 := pkg.NextUint8()
	fmt.Sprintf("random number: %v\n", n8)
}
