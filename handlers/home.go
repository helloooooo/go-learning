package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/helloooooo/go-learning/version"
)

type versionInfo struct {
	BuildTime string `json:"buildTime"`
	Commit    string `json:"commit"`
	Release   string `json:"release"`
}

func home(w http.ResponseWriter, _ *http.Request) {
	info := versionInfo{
		version.BuildTime,
		version.Commit,
		version.Release,
	}
	body, err := json.Marshal(info)
	if err != nil {
		log.Printf("Could not  encode info data %v", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
