package repository

import "Tryscon-api/models"

type Repository interface {
	User
	Task
	Organization
}

type User interface {
	GetUserByID(uuid models.UserByUUID)
	CreateUser(input models.SignUp) (models.UserByUUID, error)
	SignIn(input models.SignIn) (models.UserByUUID, error)
	DeleteUser(uuid models.UserByUUID) error
	UpdateUser(input models.UpdateUser) (models.UpdateUser, error)
}

type Task interface {
	GetTaskByID(uuid models.TaskByUUID)
	CreateTask(input models.TaskCreate) (models.TaskCreate, error)
	IsDoneTask(input models.IsDone) error
	UpdateTask(input models.UpdateTask) (models.UpdateTask, error)
	DeleteTask(uuid models.TaskByUUID) error
}

type Organization interface {
	GetOrganizationByID(uuid models.OrganizationByUUID)
	CreateOrganization(input models.CreateOrganization) (models.CreateOrganization, error)
	GetUsersInOrganization(uuid models.OrganizationByUUID) ([]models.UsersInOrganization, error)
	GetTasksInOrganization(uuid models.OrganizationByUUID) ([]models.TasksInOrganization, error)
	GetTaskByUser(uuid models.UserByUUID) (*models.Task, error)
	DeleteTaskByUser(uuid models.UserByUUID) error
	DeleteUserByID(uuid models.UserByUUID) error
	DeleteOrganizationByID(uuid models.OrganizationByUUID) error
	UpdateOrganizationByID(uuid models.OrganizationByUUID) (*models.Organization, error)
}
