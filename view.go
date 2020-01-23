package main

import (
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
)


func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) giveStatusCode(w http.ResponseWriter, r *http.Request) {

	requestVars := mux.Vars(r)
	statusCode, err := strconv.ParseUint(requestVars["statusCode"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to determine status code"))
	}

	respondWithJSON(w, int(statusCode), struct {
		Status uint64	`json:"status"`
	}{
		Status: statusCode,
	})
}

func (a *App) slowResponse(w http.ResponseWriter, r *http.Request) {

	time.Sleep(5 * time.Second)

	respondWithJSON(w, http.StatusOK, struct {
		Message string	`json:"message"`
	}{
		Message: "Response was delayed 5 seconds",
	})
}
