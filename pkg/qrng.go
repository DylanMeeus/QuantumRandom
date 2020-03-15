package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type datatype string

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
	int8buffer  []uint8
	int16buffer []uint16
	uintbuffer  []uint
)

func NextUint8() uint8 {
	fmt.Println("getting the next number")
	queryApi(u8)
	return 0
}

func queryApi(t datatype) ([]uint, error) {
	query := fmt.Sprintf("%v?length=%d&type=%v", apibase, 10, t)
	fmt.Printf("query: %v\n", query)
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Status code: %v\n", response.StatusCode)

	resp := new(ApiResponse)

	json.NewDecoder(response.Body).Decode(resp)
	fmt.Printf("decoded: %v\n", resp)
	return nil, nil
}
