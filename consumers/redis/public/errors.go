package public

import (
	"encoding/json"
	"net/http"
)

type errJSON struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err"`
}

var ErrJSON = jsonTools{}

type jsonTools struct {
}

func (jsonTools) WriteString(w http.ResponseWriter, err string, code ...int) error {
	for _, status := range code {
		w.WriteHeader(status)
	}
	return json.NewEncoder(w).Encode(errJSON{Err: err, Ok: false})
}

func (jsonTools) Write(w http.ResponseWriter, err error, code ...int) error {
	for _, status := range code {
		w.WriteHeader(status)
	}
	return json.NewEncoder(w).Encode(errJSON{Err: err.Error(), Ok: false})
}
