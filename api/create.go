package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Metudu/oss-project/data"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	var user []data.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data.CreateUser(user)
	w.WriteHeader(http.StatusOK)
}