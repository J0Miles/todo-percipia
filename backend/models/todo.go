package models

import "time"

type Todo struct {
	Id              int
	Title           string `json:"title,omitempty"`
	Description     string `json:"description,omitempty"`
	Created_at      time.Time
	Updated_at      time.Time
	Completed       bool
	EditModeEnabled bool
}
