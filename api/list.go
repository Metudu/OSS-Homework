package api

import (
	"encoding/json"
	"net/http"

	"github.com/Metudu/oss-project/data"
)

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data.UserList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}