package handler

import (
	"encoding/json"
	"net/http"

	"github.com/frodonLD/GoLangRESTAPIWithMux/model"
	"github.com/gorilla/mux"
)

func GetAllNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Notifications)
}

func GetNotificationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	for _, item := range model.Notifications {
		if item.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
