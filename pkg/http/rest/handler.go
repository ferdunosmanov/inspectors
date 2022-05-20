package rest

import (
	"github.com/ferdunosmanov/inspectors/pkg/reading"
	"github.com/gorilla/mux"
)

func InitHandlers(rs reading.Service, registrations AllRegistrations.Service) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/api/products", GetAllProductNames(rs)).Methods("GET")
	router.HandleFunc("/api/registrations", AllRegistrations(registrations)).Methods("GET")

	return router
}
