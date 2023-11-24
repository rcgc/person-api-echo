package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestPerson_Create_wrong_structure(t *testing.T) {
	data := []byte(`{"name": 123, "age": 18}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)

	p := newPerson(&StorageMockOk{})
	err := p.create(ctx)
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo %v", err)
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("Código de estado, se esperaba %d, se obtuvo %d", http.StatusBadRequest, w.Code)
	}

	resp := response{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("no se pudo hacer unmarshal al body %v", err)
	}

	wantMessage := "La persona no tiene una estructura correcta"
	if resp.Message != wantMessage {
		t.Errorf("La respuesta no es la esperada, se obtuvo %q, se esperaba %q", resp.Message, wantMessage)
	}
	
	//t.Logf("Body: %v", w.Body.String())
}

func TestPerson_Create_wrong_storage(t *testing.T) {
	data := []byte(`{"name": "RC", "age": 18}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)

	p := newPerson(&StorageMockError{})
	err := p.create(ctx)
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Código de estado, se esperaba %d, se obtuvo %d", http.StatusInternalServerError, w.Code)
	}

	resp := response{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("no se pudo hacer unmarshal al body %v", err)
	}

	wantMessage := "Hubo un problema al crear la persona"
	if resp.Message != wantMessage {
		t.Errorf("La respuesta no es la esperada, se obtuvo %q, se esperaba %q", resp.Message, wantMessage)
	}
}

func TestPerson_Create(t *testing.T) {
	data := []byte(`{"name": "RC", "age": 18}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)

	p := newPerson(&StorageMockOk{})
	err := p.create(ctx)
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo %v", err)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("Código de estado, se esperaba %d, se obtuvo %d", http.StatusCreated, w.Code)
	}

	resp := response{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("no se pudo hacer unmarshal al body %v", err)
	}

	wantMessage := "Persona creada correctamente"
	if resp.Message != wantMessage {
		t.Errorf("La respuesta no es la esperada, se obtuvo %q, se esperaba %q", resp.Message, wantMessage)
	}
}