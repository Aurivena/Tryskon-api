package models

import (
	"Tryscon-api/pkg/enum"
	"github.com/google/uuid"
)

type Task struct {
	TaskByUUID
	UpdateTask
	IsDone       bool               `json:"is_done"`
	Organization OrganizationByUUID `json:"organization"`
}

type TaskCreate struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Author      UserByUUID `json:"author"`
}

type UpdateTask struct {
	TaskCreate
	Complexity enum.Complexity `json:"complexity"`
	Status     enum.Status     `json:"status"`
}

type TaskByUUID struct {
	UUID uuid.UUID `json:"uuid"`
}

type IsDone struct {
	IsDone bool `json:"is_done"`
}
