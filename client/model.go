package client

// Login .
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GeneralResponse .
type GeneralResponse struct {
	MessageType string `json:"message_type"`
	Message string `json:"message"`
}

// LoginResponse .
type LoginResponse struct {
	GeneralResponse
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

// Community .
type Community struct {
	Name string `json:"name"`
}

// Communities slice de comunidades
type Communities []Community

// Person estructura de una persona
type Person struct {
	Name string `json:"name"`
	Age uint8	`json:"age"`
	Communities Communities `json:"communities"`
}

type Persons []Person