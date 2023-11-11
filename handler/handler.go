package handler

import "github.com/rcgc/person-api-http-net/model"

// Storage interface to be implemented by all the handlers that need to interact with memory
type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}