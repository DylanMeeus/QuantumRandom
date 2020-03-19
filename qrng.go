package pkg

import (
	"fmt"
)

type datatype string

const (
	u8  datatype = "uint8"
	u16 datatype = "uint16"
)

type cache struct {
	typ       datatype
	size      int
	ptr       int
	data      []uint
	populated bool
}

// isExhausted is true if all values of the cache have been used
func (c *cache) isExhausted() bool {
	return c.ptr == len(c.data)
}

// isEmpty is only true if this cache has not been used before
func (c *cache) isEmpty() bool {
	return !c.populated
}

// next returns the next number in the cache
func (c *cache) next() uint {
	val := c.data[c.ptr]
	c.ptr++
	return val
}

// next overrides the curent cache with new values
func (c *cache) reset(is []uint) {
	c.data = make([]uint, len(is))
	for i, v := range is {
		c.data[i] = v
	}
	c.ptr = 0
	c.populated = true
}

var (
	uint8cache  cache
	uint16cache cache
)

func init() {
	uint8cache = cache{typ: u8}
	uint16cache = cache{typ: u16}
}

// NextInt returns an int value out of convenience. The QRNG server does not natively support int
// values, thus all values will be strictly in the uint16 range.
// An error will return -1 along with the error
func NextInt() (int, error) {
	ui, err := NextUint16()
	if err != nil {
		return -1, err
	}
	return int(ui), nil
}

// NextIntN returns a batch of N uncached integers.
func NextIntN(amount int) ([]int, error) {
	uis, err := queryApi(u16, amount)
	if err != nil {
		return nil, err
	}
	is := make([]int, len(uis))
	for index, v := range uis {
		is[index] = int(v)
	}
	return is, nil
}

// NextUint8 will return the next uint8 number. If the cache is empty, it will repopulate it from
// the anu.edu servers.
func NextUint8() (uint8, error) {
	fmt.Println("getting the next number..")
	if uint8cache.isEmpty() || uint8cache.isExhausted() {
		numbers, err := queryDefaultSize(u8)
		if err != nil {
			return 0, err
		}
		uint8cache.reset(numbers)
	}
	return uint8(uint8cache.next()), nil
}

// NextUint16 will return the next uint16 number. If the cache is empty, it will repopulate it from
// the anu.edu servers.
func NextUint16() (uint16, error) {
	if uint16cache.isEmpty() || uint16cache.isExhausted() {
		numbers, err := queryDefaultSize(u16)
		if err != nil {
			return 0, err
		}
		uint16cache.reset(numbers)
	}
	return uint16(uint16cache.next()), nil
}
