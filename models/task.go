package models

import (
	"Tryscon-api/pkg/enum"
	"github.com/google/uuid"
)

type Task struct {
	TaskByID
	UpdateTask
	IsDone bool `json:"is_done"`
}

type UpdateTask struct {
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	Complexity   enum.Complexity    `json:"complexity"`
	Status       enum.Status        `json:"status"`
	Author       UserByUUID         `json:"author"`
	Organization OrganizationByUUID `json:"organization"`
}

type TaskByID struct {
	UUID uuid.UUID `json:"uuid"`
}
