package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// queryDefaultSize will request ten numbers from the API.
func queryDefaultSize(dt datatype) ([]uint, error) {
	return queryApi(dt, 10)
}

// querySingleValue will just request one number from the API.
func querySingleValue(dt datatype) (uint, error) {
	resp, err := queryApi(dt, 1)
	if err != nil {
		return 0, err
	}
	return resp[0], nil
}

// queryApi performs the actual request against the anu.edu servers.
func queryApi(dt datatype, size int) ([]uint, error) {
	query := fmt.Sprintf("%v?length=%d&type=%v", apibase, size, dt)
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	resp := new(ApiResponse)
	json.NewDecoder(response.Body).Decode(resp)
	return resp.Data, nil
}
