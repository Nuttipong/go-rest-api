package controllers

import (
	"net/http"
)

// LivenessController to check the liveness api
type LivenessController struct{}

// Get will return OK
func (c *LivenessController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("OK"))
}
