package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type datatype string

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

const (
	apibase          = "https://qrng.anu.edu.au/API/jsonI.php"
	u8      datatype = "uint8"
	u16     datatype = "uint16"
)

type ApiResponse struct {
	Length int    `json:"length"`
	Data   []uint `json:"data"`
}

var (
	uint8cache  cache
	uint16cache cache
)

func init() {
	uint8cache = cache{typ: u8}
	uint16cache = cache{typ: u16}
}

// NextUint8 will return the next uint8 number. If the cache is empty, it will repopulate it from
// the anu.edu servers.
func NextUint8() (uint8, error) {
	fmt.Println("getting the next number..")
	if uint8cache.isEmpty() || uint8cache.isExhausted() {
		numbers, err := queryApi(u8)
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
		numbers, err := queryApi(u16)
		if err != nil {
			return 0, err
		}
		uint16cache.reset(numbers)
	}
	return uint16(uint16cache.next()), nil
}

func queryApi(dt datatype) ([]uint, error) {
	query := fmt.Sprintf("%v?length=%d&type=%v", apibase, 10, dt)
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	resp := new(ApiResponse)
	json.NewDecoder(response.Body).Decode(resp)
	return resp.Data, nil
}
