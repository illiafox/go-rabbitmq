package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type getJSON struct {
	Success bool `json:"success"`

	Error struct {
		Type string `json:"type"`
		Info string `json:"info"`
	} `json:"error"`

	Quotes map[string]json.Number `json:"quotes"`
}

func Parse(key string) (map[string]string, error) {
	resp, err := http.Get("http://api.currencylayer.com/live?access_key=" + key)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	var get getJSON

	err = json.NewDecoder(resp.Body).Decode(&get)

	if err != nil {
		return nil, fmt.Errorf("decoding json: %w", err)
	}

	if !get.Success {
		return nil, fmt.Errorf("api error: %s", get.Error.Type+":"+get.Error.Info)
	}

	var m = map[string]string{}

	for k, v := range get.Quotes {
		m[k] = v.String()
	}

	return m, nil
}
