package models

import "github.com/google/uuid"

type Organization struct {
	OrganizationByUUID
	UpdateOrganization
	Users []UserByUUID `json:"users"`
	Tasks []Task       `json:"tasks"`
}

type UpdateOrganization struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type OrganizationByUUID struct {
	UUID uuid.UUID `json:"uuid"`
}
