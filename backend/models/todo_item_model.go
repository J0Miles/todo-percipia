package models

type TodoItemModel struct {
	Id          int `gorm:primary_key`
	Title       string
	Description string
	Created_at  string
	Updated_at  string
	Completed   bool
}
