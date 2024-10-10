package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func failResponseWriter(w http.ResponseWriter, err error, errStatusCode int) {
	w.Header().Set("Content-Type", "application/json")

	var response Response
	w.WriteHeader(errStatusCode)
	response.StatusCode = errStatusCode
	response.Message = err.Error()
	response.Data = nil

	responseBytes, _ := json.Marshal(response)
	w.Write(responseBytes)
}

func successResponseWriter(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	var response Response
	w.WriteHeader(http.StatusOK)
	response.StatusCode = http.StatusOK
	response.Message = "Success"
	response.Data = data

	responseBytes, _ := json.Marshal(response)
	w.Write(responseBytes)
}

func WriteResponse(w http.ResponseWriter, err error, data any) {
	switch err.(type) {
	case nil:
		successResponseWriter(w, data)
	case *ErrBadRequest:
		failResponseWriter(w, err, http.StatusBadRequest)
	case *ErrNotFound:
		failResponseWriter(w, err, http.StatusNotFound)
	case *ErrInternalServerError:
		failResponseWriter(w, err, http.StatusInternalServerError)
	case *ErrUnauthorized:
		failResponseWriter(w, err, http.StatusUnauthorized)
	case *ErrForbidden:
		failResponseWriter(w, err, http.StatusForbidden)
	default:
		failResponseWriter(w, err, http.StatusInternalServerError)
	}
}
