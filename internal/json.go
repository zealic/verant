package internal

import (
	"encoding/json"
	"net/http"
)

//UnmarshalJSON from url unmarshal json to struct
func UnmarshalJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
