package storage

import (
	"testing"

	"github.com/rcgc/person-api-echo/model"
)

func TestCreate(t *testing.T) {
	table := []struct{
		name string
		person *model.Person
		wantError error
	}{
		{"Empty person", nil, model.ErrPersonCanNotBeNil},
		{"RC", &model.Person{Name: "RC"}, nil},
		{"Joe", &model.Person{Name: "Joe"}, nil},
	}

	m := NewMemory()
	for _, v := range table {
		t.Run(v.name, func( t *testing.T){
			gotError := m.Create(v.person)
			if gotError != v.wantError {
				t.Errorf("Se esperaba el error %v, se obtuvo %v", v.wantError, gotError)
			}
		})
	}

	wantCount := len(table) - 1
	if m.currentID != wantCount{
		t.Errorf("Se esperaba %d registros, se obtuvo %d", wantCount, m.currentID)
	}
}

/*
func TestCreate_empty_person(t *testing.T) {
	m := NewMemory()
	err := m.Create(nil)

	if err == nil {
		t.Error("Se esperaba error, se obtuvo nil")
	}

	if err != model.ErrPersonCanNotBeNil {
		t.Errorf("Se esperaba el error %v, se obtuvo el error %v", model.ErrPersonCanNotBeNil, err)
	}
}

func TestCreate_count_elements(t *testing.T){
	m := NewMemory()
	p := model.Person{Name: "RC"}
	err := m.Create(&p)

	if err != nil {
		t.Errorf("No se esperaba error, se obtuvo: %v", err)
	}

	if m.currentID != 1 {
		t.Errorf("Se esperaba 1 elemento, se obtuvo: %d", m.currentID)
	}
}
*/