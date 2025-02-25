// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Note struct {
	NoteID    string `json:"noteId"`
	UserID    string `json:"userId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type Query struct {
}

type Task struct {
	TaskID      string `json:"taskId"`
	UserID      string `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
}

type User struct {
	UserID string  `json:"userId"`
	Name   string  `json:"name"`
	Notes  []*Note `json:"notes"`
	Tasks  []*Task `json:"tasks"`
}
