package api

import (
	"encoding/json"
	"net/http"

	"github.com/Metudu/oss-project/data"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	type requestBody struct {
		Index int `json:"index"`
		User data.User `json:"user"`
	}

	var req requestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data.UpdateUser(req.Index, req.User)
	w.WriteHeader(http.StatusOK)
}