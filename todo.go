package todo

import "errors"

type Todo struct {
	ID          int    `json:"id "db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	ID     int
	UserID int
	ListID int
}

type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"completed"`
}

type ListsItem struct {
	ID     int
	ListID int
	ItemID int
}

type UpdateTodo struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (u *UpdateTodo) Validate() error {
	if u.Title == nil && u.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateItem struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}
