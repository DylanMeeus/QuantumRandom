# QuantumRandom

This library gets the data from the ANU Quantum Random Numbers server.
For more information, please check out their [website](http://qrng.anu.edu.au/index.php), or refer
to their papers:

- [Real time demonstration of high bitrate quantum random number generation with coherent laster
  light](https://aip.scitation.org/doi/10.1063/1.3597793)

- [Maximization of Extractable Randomness in a Quantum Random Number
  Generator](https://journals.aps.org/prapplied/abstract/10.1103/PhysRevApplied.3.054004)

# API

The full docs are on [godoc](https://godoc.org/github.com/DylanMeeus/QuantumRandom/pkg)

But the  basic functions to get started are:

- `NextInt(): int` -> return a single QRNG number
- `NextIntN(amount int): []int` -> return 'amount' of QRNG numbers
- `NextUint8() -> uint8` -> return a single QRNG uint8
- `NextUint16() -> uint16` -> return a single QRNG uint16

