package repository

import (
	todo "TodoListREST"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemQuery, listId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.completed FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.completed FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *TodoItemPostgres) Update(userId, itemId int, updateInput todo.UpdateItem) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateInput.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *updateInput.Title)
		argId++
	}

	if updateInput.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *updateInput.Description)
		argId++
	}

	if updateInput.Done != nil {
		setValues = append(setValues, fmt.Sprintf("completed=$%d", argId))
		args = append(args, *updateInput.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d",
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}
