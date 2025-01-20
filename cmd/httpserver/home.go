package main

import "net/http"

func (s *PackServer) Home(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Welcome to the packager API"))
	if err != nil {
		return
	}
}
