package main

import (
	"encoding/json"
	"net/http"
)

type DefaultPacksResponse struct {
	Packs []int `json:"packs"`
}

func (s *PackServer) DefaultPacks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(DefaultPacksResponse{Packs: defaultPackSizes})
}
