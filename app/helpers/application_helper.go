package helpers

import (
	"encoding/json"
	"net/http"
)

func RenderJson(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data = map[string]string{"errors": err.Error()}
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(data)
}
