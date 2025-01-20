package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type MinPackRequest struct {
	Order     int   `json:"order"`
	PackSizes []int `json:"packaging_sizes"`
}

var (
	defaultPackSizes = []int{250, 500, 1000, 2000, 5000}
)

func (s *PackServer) Packs(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, "could not close request body", http.StatusInternalServerError)
			return
		}
	}(r.Body)

	var minPackRequest MinPackRequest
	err = json.Unmarshal(body, &minPackRequest)

	log.Println("minPackRequest", minPackRequest)
	if err != nil {
		http.Error(w, "could not decode request body", http.StatusBadRequest)
		return
	}

	if minPackRequest.Order <= 0 {
		http.Error(w, "order must be greater than 0", http.StatusBadRequest)
		return

	}

	if len(minPackRequest.PackSizes) == 0 {
		minPackRequest.PackSizes = defaultPackSizes
	}

	packs := s.packer.CalculateMinPacks(minPackRequest.Order, minPackRequest.PackSizes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(packs)
}
