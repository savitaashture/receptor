package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cloudfoundry-incubator/receptor"
)

func writeUnknownErrorResponse(w http.ResponseWriter, err error) error {
	return writeJSONResponse(w, http.StatusInternalServerError, receptor.Error{
		Type:    receptor.UnknownError,
		Message: err.Error(),
	})
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, jsonObj interface{}) error {
	jsonBytes, err := json.Marshal(jsonObj)
	if err != nil {
		panic("Unable to encode JSON: " + err.Error())
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(jsonBytes)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err = w.Write(jsonBytes)
	return err
}