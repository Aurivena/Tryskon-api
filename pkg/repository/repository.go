package repository

import "Tryscon-api/models"

type Repository interface {
	GetUserByID(id int)
	CreateUser(input models.SignUp) (models.SignUp, error)
}
