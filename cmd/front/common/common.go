package common

import "net/http"

func SendBadRequestResponse(w http.ResponseWriter, errString string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(errString))
}
