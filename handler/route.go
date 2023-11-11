package handler

import (
	"net/http"

	"github.com/rcgc/person-api-http-net/middleware"
)

// RouterPerson .
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", middleware.Log(h.create))
	mux.HandleFunc("/v1/persons/update", middleware.Log(h.update))
	mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
	mux.HandleFunc("/v1/persons/get-by-id", middleware.Log(h.getByID))
	mux.HandleFunc("/v1/persons/get-all", middleware.Log(middleware.Authentication(h.getAll)))
}

// RouteLogin .
func RouteLogin(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)

	mux.HandleFunc("/v1/login", h.login)
}