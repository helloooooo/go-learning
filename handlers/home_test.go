package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	buildtime := time.Now().Format("20060102_03:04:05")
	commit := os.Getenv("COMMIT")
	release := os.Getenv("RELEASE")
	home(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Statu code is wrong. Have: %d, want:%d", have, want)
	}

	res, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	var info versionInfo
	err = json.Unmarshal(res, &info)
	if err != nil {
		t.Fatal(err)
	}
	if info.BuildTime != buildtime || info.Commit != commit || info.Release != release {
		t.Errorf("response body is wrong. json = %v", info)
	}
}
