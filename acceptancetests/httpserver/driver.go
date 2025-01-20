package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/vcayetano/fulfillment/specifications"
	"io"
	"log"
	"net/http"
	"strings"
)

type Driver struct {
	BaseURL string
	Client  *http.Client
}

func (d Driver) CalculateMinPacks(orderSize int, packageSizes []int) specifications.Packages {

	type MinPackRequest struct {
		Order     int   `json:"order"`
		PackSizes []int `json:"packaging_sizes"`
	}

	requestBody, _ := json.Marshal(MinPackRequest{
		Order:     orderSize,
		PackSizes: packageSizes,
	})

	res, err := d.Client.Post(d.BaseURL+"/packs", "application/json", strings.NewReader(string(requestBody)))

	if err != nil {
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(res.Body)

	result := specifications.Packages{}
	packages, err := io.ReadAll(res.Body)

	fmt.Println("packages: ", string(packages))

	if err != nil {
		return nil
	}

	err = json.Unmarshal(packages, &result)

	if err != nil {
		return nil
	}

	return result
}
