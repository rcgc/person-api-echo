package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func loginClient(url, email, password string) LoginResponse{
	login := Login{
		Email:    email,
		Password: password,
	}

	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("Error en marshal de login: %v", err)
	}

	resp := httpClient(http.MethodPost, url, "", data)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Se esperaba código 200, se obtuvo: %d, respuesta: %s", resp.StatusCode, string(body))
	}

	if err != nil {
		log.Fatalf("Error en lectura del body en login: %v", err)
	}

	dataResponse := LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error en el unmarshal del body en login: %v", err)
	}

	return dataResponse
}