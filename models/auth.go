package models

import (
	"Tryscon-api/pkg/enum"
	"github.com/google/uuid"
)

type SignIn struct {
	Login    string `json:"login"`
	Password string `json:"password_hash"`
	Email    string `json:"email"`
}

type SignUp struct {
	UserByUUID
	SignIn
	Status enum.Status `json:"status"`
}

type UpdateUser struct {
	SignIn
	Description string `json:"description"`
}

type UserByUUID struct {
	UUID uuid.UUID `json:"uuid"`
}

type User struct {
	UserByUUID
	SignIn
}
