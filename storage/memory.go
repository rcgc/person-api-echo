package storage

import (
	"fmt"

	"github.com/rcgc/person-api-http-net/model"
)

// Memory is a map of Persons.
type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

// NewMemory returns an instance of Memory.
func NewMemory() Memory{
	persons := make(map[int]model.Person)

	return Memory{
		currentID: 0,
		Persons: persons,
	}
}

// Create adds a new person to map
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	m.currentID++
	m.Persons[m.currentID] = *person
	return nil
}

// Update actualiza una persona en el slice de memoria
func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	m.Persons[ID] = *person

	return nil
}

// Delete borra de la memoria la persona
func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	delete(m.Persons, ID)

	return nil
}

// GetByID retorna una persona por el ID
func (m *Memory) GetByID(ID int) (model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	return person, nil
}

// GetAll retorna todas las persona que están en la memoria
func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}

	return result, nil
}