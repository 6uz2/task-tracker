package dto

import "time"

type TaskCollection struct {
	LastTaskId int
	Tasks      []TaskProperties
}

type TaskProperties struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
