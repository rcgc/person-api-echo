package handler

import (
	"github.com/labstack/echo"
	"github.com/rcgc/person-api-echo/middleware"
)

// RouterPerson .
func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)
	person := e.Group("/v1/persons")
	person.Use(middleware.Authentication)

	person.POST("", h.create)
	person.PUT("/:id", h.update)
	person.DELETE("/:id", h.delete)
	person.GET("/:id", h.getByID) // in case of similar routes, the last one overwrites the previous one
	person.GET("", h.getAll)
}

// RouteLogin .
func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/v1/login", h.login)
}