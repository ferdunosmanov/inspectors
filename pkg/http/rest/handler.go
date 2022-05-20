package rest

import (
	"github.com/ferdunosmanov/inspectors/pkg/GetAllRegistrations"
	"github.com/ferdunosmanov/inspectors/pkg/reading"
	"github.com/gorilla/mux"
)

func InitHandlers(rs reading.Service, registrations GetAllRegistrations.Service) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/api/products", GetAllProductNames(rs)).Methods("GET")
	router.HandleFunc("/api/registrations", GetAllRegistrations(registrations)).Methods("GET")

	return router
}
