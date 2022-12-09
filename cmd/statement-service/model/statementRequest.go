package model

type Request struct {
	Email string `json:"email" validate:"email"`
}
