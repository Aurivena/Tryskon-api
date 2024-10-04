package repository

type Repository interface {
	GetUserByID(id int)
}
