package api

import (
	"net/http"
	"strconv"

	"github.com/Metudu/oss-project/data"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	index, err := strconv.Atoi(r.URL.Query().Get("index"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if index < 0 || index >= len(data.UserList) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	data.DeleteUser(index)
	w.WriteHeader(http.StatusOK)
}