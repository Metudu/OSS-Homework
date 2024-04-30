package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Metudu/oss-project/data"
)

func Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")	

	index, err := strconv.Atoi(r.URL.Query().Get("index"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := data.ReadUser(index)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}