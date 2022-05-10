package rest

import (
	"encoding/json"
	"net/http"

	"github.com/ferdunosmanov/inspectors/pkg/reading"
)

func welcomeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to our Inspectors App")
	}
}

func GetAllProductNames(rs reading.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cs, err := rs.GetAllProductNames()
		if err != nil {
			http.Error(w, "Cannot Process your request at this time....", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(cs)
	}
}
