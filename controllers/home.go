package controllers

import (
	"fmt"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you are in the %s endpoint", r.URL.Path)
}
