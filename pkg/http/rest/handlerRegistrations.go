package rest

import (
	"encoding/json"
	"net/http"
)

func GetAllRegistrations(registrations GetAllRegistrations.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ir, err := registrations.GetAllRegistrations()
		if err != nil {
			http.Error(w, "Cannot Process your request at this time....", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(ir)
	}
}
