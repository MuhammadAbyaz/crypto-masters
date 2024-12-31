package api

import (
	"cryptomasters/structures"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*structures.Rate, error) {
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	var response CEXResponse
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		dataStream, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(dataStream, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	return &structures.Rate{Currency: currency, Price: float64(response.Bid)}, nil
}
