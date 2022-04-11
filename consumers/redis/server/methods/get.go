package methods

import (
	"consumer/database"
	"consumer/public"
	"encoding/json"
	"fmt"
	"net/http"
)

type Methods struct {
	Redis *database.Redis
}

func (m Methods) Update(w http.ResponseWriter, r *http.Request) {
	var get getJSON

	err := json.NewDecoder(r.Body).Decode(&get)
	if err != nil {
		public.ErrJSON.Write(w, fmt.Errorf("json decoding: %w", err), http.StatusForbidden)
		return
	}

	if get.Base == "" {
		public.ErrJSON.WriteString(w, "'base' field is empty", http.StatusForbidden)
		return
	}

	price, err := m.Redis.Currencies.Get(get.Base)
	if err != nil {
		public.ErrJSON.Write(w, fmt.Errorf("get currency: %w", err), http.StatusForbidden)
		return
	}

	if price == "" {
		public.ErrJSON.WriteString(w, "currency not found", http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(retJSON{
		Ok:       true,
		getJSON:  &get,
		Currency: []byte(price),
	})
}

type retJSON struct {
	Ok bool `json:"ok"`
	*getJSON
	Currency json.RawMessage `json:"currency"`
}

type getJSON struct {
	Base string `json:"base"`
}
