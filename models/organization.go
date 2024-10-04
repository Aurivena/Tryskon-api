package models

import "github.com/google/uuid"

type Organization struct {
	CreateOrganization
	UsersInOrganization
	TasksInOrganization
}

type UsersInOrganization struct {
	Users []UserByUUID `json:"users"`
}

type TasksInOrganization struct {
	Tasks []Task `json:"tasks"`
}

type UpdateOrganization struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type CreateOrganization struct {
	OrganizationByUUID
	UpdateOrganization
}

type OrganizationByUUID struct {
	UUID uuid.UUID `json:"uuid"`
}
