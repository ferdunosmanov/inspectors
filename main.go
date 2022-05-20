package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ferdunosmanov/inspectors/pkg/GetAllRegistrations"
	"github.com/ferdunosmanov/inspectors/pkg/http/rest"
	"github.com/ferdunosmanov/inspectors/pkg/reading"
	"github.com/ferdunosmanov/inspectors/pkg/storage"
)

func main() {

	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("error while setting up storage:", err)
	}

	rs := reading.NewService(r)
	registrations := GetAllRegistrations.NewService(r)

	fmt.Println("Starting server on port: 8080...")
	router := rest.InitHandlers(rs, registrations)
	log.Fatal(http.ListenAndServe(":8080", router))
}
