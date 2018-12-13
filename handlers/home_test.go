package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	home(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Statu code is wrong. Have: %d, want:%d", have, want)
	}

	grreting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if have, want := string(grreting), "Hello! Your request was processed."; have != want {
		t.Errorf("The greeting is wrong. Have:%s, want:%s", have, want)
	}
}
