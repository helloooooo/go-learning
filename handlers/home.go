package handlers

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, _ *htto.Request) {
	fmt.Fprint(w, "Hello! Your request was processed")
}